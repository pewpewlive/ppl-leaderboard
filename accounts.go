package leaderboard

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/pewpewlive/common-go/ppl_types"
)

// Maps account IDs to names
func GetAccountMapFromInfo(accounts []ppl_types.AccountInfo) map[string]string {
	data := map[string]string{}

	for _, account := range accounts {
		data[account.AccountID] = account.Username
	}

	return data
}

// Maps account IDs to names
func GetAccountMapFromCSV(path string) map[string]string {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil
	}

	data, err := gocsv.CSVToMap(file)
	if err != nil {
		return nil
	}

	return data
}
