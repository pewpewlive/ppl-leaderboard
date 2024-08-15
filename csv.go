package leaderboard

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/pewpewlive/common-go/ppl_types"
)

// Utility to read `[]ppl_types.HofEntry` from csv file.
func GetScoresFromCSV(path string) ([]ppl_types.HofEntry, error) {
	var scores []ppl_types.HofEntry

	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return scores, err
	}
	err = gocsv.UnmarshalFile(file, &scores)
	if err != nil {
		return scores, err
	}
	return scores, nil
}

// Utility to read `[]ppl_types.AccountInfo` from csv file.
func GetAccountsFromCSV(path string) ([]ppl_types.AccountInfo, error) {
	var accounts []ppl_types.AccountInfo
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return accounts, err
	}
	err = gocsv.UnmarshalFile(file, &accounts)
	if err != nil {
		return accounts, err
	}
	return accounts, nil
}
