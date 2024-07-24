package leaderboard

import (
	"math"
	"sort"
	"time"

	"github.com/pewpewlive/common-go/helpers"
	"github.com/pewpewlive/common-go/ppl_json"
)

type Score struct {
	PlayerAccountIDs []string
	Value            int64
	Date             time.Time
	Country          string
}

type LevelData struct {
	Scores1p []Score
	Scores2p []Score
}

type PlayerScore struct {
	AccountID        string
	AccumulatedScore float64
	Country          string
	NumberOfWRs      int
}

func ComputePlayerScores(levelData map[ppl_json.LevelFullID]LevelData) []PlayerScore {
	type PlayerTempScore struct {
		AccumulatedScore     float64
		AccumulatedCountries []string
		NumberOfWRs          int
	}

	temp := map[string]PlayerTempScore{}

	for _, data := range levelData {
		numerator := math.Pow(float64(len(data.Scores1p)), 1.0/6.0) * 100.0

		rank := 1
		for _, score := range data.Scores1p {
			// performance = (total_scores^(1/6)) * 100 / (((rank+1) / 2) ^ (1/2))
			denominator := math.Sqrt(float64(rank) / 2.0)
			delta := numerator / denominator
			var pData = temp[score.PlayerAccountIDs[0]]
			if rank == 1 {
				pData.NumberOfWRs++
			}
			pData.AccumulatedScore += delta
			pData.AccumulatedCountries = append(pData.AccumulatedCountries, score.Country)
			temp[score.PlayerAccountIDs[0]] = pData
			rank++
		}
	}

	output := make([]PlayerScore, 0)
	for k, v := range temp {
		country := ""
		if len(v.AccumulatedCountries) == 1 {
			country = v.AccumulatedCountries[0]
		} else {
			country = helpers.MostFrequentString(v.AccumulatedCountries)
		}
		output = append(output, PlayerScore{k, v.AccumulatedScore, country, v.NumberOfWRs})
	}

	sort.SliceStable(output, func(i, j int) bool {
		return output[i].AccumulatedScore > output[j].AccumulatedScore
	})

	return output
}
