package leaderboard

import (
	"sort"
	"time"

	"github.com/pewpewlive/common-go/ppl_types"
)

type Score struct {
	PlayerAccountIDs []string
	// Either a score, or a number of frames
	Value int64
	// 0 = score. 1 = speed run.
	Type    int32
	Date    time.Time
	Country string
}

type PlayerRank struct {
	AccountID        string
	AccumulatedScore float64
	Country          string
	NumberOfWRs      int
}

func ComputePlayerScores(levelData []ppl_types.HofEntry) []PlayerRank {
	type PlayerTempScore struct {
		AccumulatedScore     float64
		AccumulatedCountries []string
		NumberOfWRs          int
	}

	// scores1p := slices.DeleteFunc(levelData, func(level ppl_types.HofEntry) bool {
	// 	return len(level.PlayerAccountIDs) != 1 || level.ValueType == 1
	// })
	// scores2p := slices.DeleteFunc(levelData, func(level ppl_types.HofEntry) bool {
	// 	return len(level.PlayerAccountIDs) != 2 || level.ValueType == 1
	// })
	// speedruns1p := slices.DeleteFunc(levelData, func(level ppl_types.HofEntry) bool {
	// 	return len(level.PlayerAccountIDs) != 1 || level.ValueType == 0
	// })
	// speedruns2p := slices.DeleteFunc(levelData, func(level ppl_types.HofEntry) bool {
	// 	return len(level.PlayerAccountIDs) != 2 || level.ValueType == 0
	// })

	// temp := map[string]PlayerTempScore{}

	// for _, data := range scores1p {
	// 	numerator := math.Pow(float64(len(scores1p)), 1.0/6.0) * 100.0

	// 	rank := 1
	// 	for _, score := range data.Scores1p {
	// 		// performance = (total_scores^(1/6)) * 100 / (((rank+1) / 2) ^ (1/2))
	// 		denominator := math.Sqrt(float64(rank) / 2.0)
	// 		delta := numerator / denominator
	// 		pData := temp[score.PlayerAccountIDs[0]]
	// 		if rank == 1 {
	// 			pData.NumberOfWRs++
	// 		}
	// 		pData.AccumulatedScore += delta
	// 		pData.AccumulatedCountries = append(pData.AccumulatedCountries, score.Country)
	// 		temp[score.PlayerAccountIDs[0]] = pData
	// 		rank++
	// 	}
	// }

	output := make([]PlayerRank, 0)
	// for k, v := range temp {
	// 	output = append(output, PlayerRank{
	// 		AccountID:        k,
	// 		AccumulatedScore: v.AccumulatedScore,
	// 		Country:          helpers.MostFrequentString(v.AccumulatedCountries),
	// 		NumberOfWRs:      v.NumberOfWRs,
	// 	})
	// }

	sort.SliceStable(output, func(i, j int) bool {
		return output[i].AccumulatedScore > output[j].AccumulatedScore
	})

	return output
}
