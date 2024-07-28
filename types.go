package leaderboard

import (
	"github.com/pewpewlive/common-go/ppl_types"
)

type LevelLeaderboard struct {
	LevelUUID string
	// 0 = score. 1 = speedrun.
	LeaderboardType int32
	PlayerCount     int32
	Scores          []ppl_types.HofEntry
}

type tempPlayerRank struct {
	AccumulatedScore     float64
	AccumulatedCountries []string
	NumberOfWRs          int
}

type PlayerRank struct {
	AccountID        string
	AccumulatedScore float64
	Country          string
	NumberOfWRs      int
}
