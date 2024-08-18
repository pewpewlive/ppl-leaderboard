package main

import (
	"fmt"
	"log"

	"github.com/pewpewlive/common-go/helpers"
	leaderboard "github.com/pewpewlive/ppl-leaderboard"
)

func main() {
	scores, scoresErr := leaderboard.GetScoresFromCSV("data/score_data.csv")
	if scoresErr != nil {
		log.Fatal(scoresErr)
	}
	accounts, accountsErr := leaderboard.GetAccountsFromCSV("data/account_data.csv")
	if accountsErr != nil {
		log.Fatal(accountsErr)
	}

	accountMap := map[string]string{}
	for _, account := range accounts {
		accountMap[account.AccountID] = account.Username
	}

	leaderboards := leaderboard.GetLeaderboardsFromScores(scores, accounts)
	ranks := leaderboard.ComputePlayerRanks(leaderboards)

	fmt.Printf("The top 10 players:\n")
	for i, rank := range ranks[:10] {
		accountName := helpers.StripColorsFromString(accountMap[rank.AccountID])
		fmt.Printf("%d. %s (%s) - Score: %.2f, WRs: %d\n", i+1, accountName, rank.AccountID, rank.AccumulatedScore, rank.NumberOfWRs)
	}
}
