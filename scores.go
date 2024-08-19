package leaderboard

import (
	"encoding/json"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/pewpewlive/common-go/helpers"
	"github.com/pewpewlive/common-go/ppl_types"
)

var LevelInBlacklist = map[string]bool{
	// "Om8xrkBMgDfX4fuyJYVk1": true, // Pew-Man
	"HpfC42TNwH5AEmvWkkQof": true, // Simon Says
	// "ZlxFwr1dXdXsZslFNyzL_": true, // Inertiacs spawner
	// "81TpajqWIlSXYmZrfjSjw": true, // Inertiac World
	"7fa9_mUjTldAgdgDmrB4q": true, // Don't touch grass
	"xo8lychHUy7lass08Cyjb": true, // Pew Pong
}

// Sorts the scores based on their value and date
func SortScores(scores []ppl_types.HofEntry) {
	sort.SliceStable(scores, func(i, j int) bool {
		if scores[i].Value > scores[j].Value {
			return true
		}
		if scores[i].Value < scores[j].Value {
			return false
		}
		return time.Unix(scores[i].Date, 0).Before(time.Unix(scores[j].Date, 0))
	})
}

type LeaderboardKey struct {
	LevelUUID    string
	LevelVersion int32
	ValueType    int32
	PlayerCount  int32
}

// Maps player scores to a given level UUID with leaderboard type and player count
func GetLeaderboardsFromScores(scores []ppl_types.HofEntry, accounts []ppl_types.AccountInfo) []LevelLeaderboard {
	data := make([]LevelLeaderboard, 0)

	// Create a map, from account IDs to account usernames.
	accountMap := map[string]string{}
	for _, account := range accounts {
		accountMap[account.AccountID] = account.Username
	}

	leaderboards := map[LeaderboardKey][]ppl_types.HofEntry{}

	for _, score := range scores {
		accountIDs := strings.Split(score.PlayerAccountIDs, "|")

		// Account missing in the accountMap, ignore it
		accountDidntExist := false
		for _, id := range accountIDs {
			if accountMap[id] == "" {
				accountDidntExist = true
				break
			}
		}
		if accountDidntExist {
			continue
		}

		leaderboardKey := LeaderboardKey{
			LevelUUID:    score.LevelUUID,
			LevelVersion: score.LevelVersion,
			PlayerCount:  int32(len(accountIDs)),
			ValueType:    score.ValueType,
		}

		// Create a new leaderboard entry if necessary
		if _, exists := leaderboards[leaderboardKey]; !exists {
			leaderboards[leaderboardKey] = []ppl_types.HofEntry{}
		}

		// Add the score to the leaderboard
		leaderboards[leaderboardKey] = append(leaderboards[leaderboardKey], ppl_types.HofEntry{
			PlayerAccountIDs: score.PlayerAccountIDs,
			LevelUUID:        score.LevelUUID,
			LevelVersion:     score.LevelVersion,
			Value:            score.Value,
			ValueType:        score.ValueType,
			Date:             score.Date,
			Country:          score.Country,
		})
	}

	// Find highest version of every level
	highestVersionOfLevel := map[string]int32{}
	for key := range leaderboards {
		highestVersionOfLevel[key.LevelUUID] = max(highestVersionOfLevel[key.LevelUUID], key.LevelVersion)
	}

	// Delete every leaderboard associated with a level that is not at the latest version.
	for key := range leaderboards {
		if key.LevelVersion < highestVersionOfLevel[key.LevelUUID] {
			delete(leaderboards, key)
		}
	}

	// Populate the actual data
	for key, leaderboard := range leaderboards {

		data = append(data, LevelLeaderboard{
			LevelUUID:       key.LevelUUID,
			LeaderboardType: key.ValueType,
			PlayerCount:     key.PlayerCount,
			Scores:          leaderboard,
		})
	}

	// Sort the scores
	for _, leaderboard := range data {
		SortScores(leaderboard.Scores)
	}

	return data
}

// Computes the rank of each player for a given score slice
func ComputeRankForGivenScores(scores []ppl_types.HofEntry, playerRanks map[string]tempPlayerRank) {
	numerator := math.Pow(float64(len(scores)), 1.0/6.0) * 100.0

	rank := 1

	// Keep track of already counted players
	countedAccounts := map[string]bool{}
	for _, score := range scores {
		// performance = (total_scores^(1/6)) * 100 / (((rank+1) / 2) ^ (1/2))
		denominator := math.Sqrt(float64(rank) / 2.0)
		delta := numerator / denominator

		accountIDs := strings.Split(score.PlayerAccountIDs, "|")

		for _, player := range accountIDs {
			if countedAccounts[player] {
				// Skip the player's score if it is already accounted for the leaderboard
				continue
			}

			playerData := playerRanks[player]
			if rank == 1 {
				playerData.NumberOfWRs++
			}

			playerData.AccumulatedScore += delta
			playerData.AccumulatedCountries = append(playerData.AccumulatedCountries, score.Country)
			playerRanks[player] = playerData

			countedAccounts[player] = true
		}
		rank++
	}
}

// Computes the player ranks based on the score map
func ComputePlayerRanks(levelLeaderboards []LevelLeaderboard) []PlayerRank {
	playerRanks := map[string]tempPlayerRank{}

	for _, leaderboard := range levelLeaderboards {
		if LevelInBlacklist[leaderboard.LevelUUID] {
			continue
		}

		ComputeRankForGivenScores(leaderboard.Scores, playerRanks)
	}

	ranks := make([]PlayerRank, 0)
	for k, v := range playerRanks {
		ranks = append(ranks, PlayerRank{
			AccountID:        k,
			AccumulatedScore: v.AccumulatedScore,
			Country:          helpers.MostFrequentString(v.AccumulatedCountries),
			NumberOfWRs:      v.NumberOfWRs,
		})
	}

	sort.SliceStable(ranks, func(i, j int) bool {
		return ranks[i].AccumulatedScore > ranks[j].AccumulatedScore
	})

	return ranks
}

// Returns a jsonified version of the player ranks
func JsonifyRanks(playerRanks []PlayerRank, accountMap map[string]string) (string, error) {
	ranks := []interface{}{}
	for _, v := range playerRanks {
		nickname := accountMap[v.AccountID]

		if nickname == "" { // Account was deleted
			continue
		}

		ranks = append(ranks, v.AccountID, nickname, math.Round(v.AccumulatedScore*100)/100, v.Country, v.NumberOfWRs)
	}

	jsonStr, err := json.Marshal(ranks)
	if err != nil {
		return "", err
	}

	return string(jsonStr), nil
}
