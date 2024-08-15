package main

import (
	"fmt"
	"log"

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

	leaderboards := leaderboard.GetLeaderboardsFromScores(scores, accounts)
	ranks := leaderboard.ComputePlayerRanks(leaderboards)

	fmt.Printf("The account IDs of the top 10 ranks:\n")
	for i, rank := range ranks {
		fmt.Printf("%d. %s\n", i, rank.AccountID)
		if i > 10 {
			return
		}
	}
}
