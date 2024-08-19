package leaderboard

import (
	"testing"

	"github.com/pewpewlive/common-go/ppl_types"
	"github.com/stretchr/testify/assert"
)

func TestSortScores(t *testing.T) {
	entry1 := ppl_types.HofEntry{
		PlayerAccountIDs: "accid_1",
		Value:            300,
		Date:             1000,
	}
	entry2 := ppl_types.HofEntry{
		PlayerAccountIDs: "accid_2",
		Value:            200,
		Date:             1000,
	}
	entry3 := ppl_types.HofEntry{
		PlayerAccountIDs: "accid_2",
		Value:            200,
		Date:             1001,
	}
	entry4 := ppl_types.HofEntry{
		PlayerAccountIDs: "accid_3",
		Value:            200,
		Date:             1002,
	}
	entry5 := ppl_types.HofEntry{
		PlayerAccountIDs: "accid_3",
		Value:            100,
		Date:             1000,
	}
	unsortedScores := []ppl_types.HofEntry{
		entry5,
		entry2,
		entry4,
		entry3,
		entry1,
	}
	expectedSortedScores := []ppl_types.HofEntry{
		entry1,
		entry2,
		entry3,
		entry4,
		entry5,
	}
	SortScores(unsortedScores)
	assert.Equal(t, expectedSortedScores, unsortedScores)
}

func TestGetLeaderboardsFromScores_MultipleLevelVersion(t *testing.T) {
	// Test that only the latest version of a level is kept
	scores := []ppl_types.HofEntry{
		{PlayerAccountIDs: "foo", LevelUUID: "level1", LevelVersion: 1, Value: 10, ValueType: 0},
		{PlayerAccountIDs: "bar", LevelUUID: "level1", LevelVersion: 3, Value: 30, ValueType: 0},
		{PlayerAccountIDs: "qux", LevelUUID: "level1", LevelVersion: 2, Value: 20, ValueType: 0},
	}
	accounts := []ppl_types.AccountInfo{
		{AccountID: "bar", Username: "Bar"},
	}
	leaderboards := GetLeaderboardsFromScores(scores, accounts)
	expectedLeaderboard := []LevelLeaderboard{
		{
			LevelUUID:       "level1",
			LeaderboardType: 0,
			PlayerCount:     1,
			Scores:          []ppl_types.HofEntry{{PlayerAccountIDs: "bar", LevelUUID: "level1", LevelVersion: 3, Value: 30, ValueType: 0}},
		},
	}
	assert.Equal(t, expectedLeaderboard, leaderboards)
}

func TestGetLeaderboardsFromScores_AccountsDeleted(t *testing.T) {
	// Test that accounts that don't exist anymore are ignored
	scores := []ppl_types.HofEntry{
		{PlayerAccountIDs: "foo", LevelUUID: "level1", Value: 10, ValueType: 0},
		{PlayerAccountIDs: "bar", LevelUUID: "level1", Value: 20, ValueType: 0},
		{PlayerAccountIDs: "qux", LevelUUID: "level1", Value: 30, ValueType: 0},
	}
	accounts := []ppl_types.AccountInfo{
		{AccountID: "foo", Username: "Foo"},
		{AccountID: "qux", Username: "Qux"},
	}
	leaderboards := GetLeaderboardsFromScores(scores, accounts)
	expectedLeaderboard := []LevelLeaderboard{
		{
			LevelUUID:       "level1",
			LeaderboardType: 0,
			PlayerCount:     1,
			Scores: []ppl_types.HofEntry{{PlayerAccountIDs: "qux", LevelUUID: "level1", Value: 30, ValueType: 0},
				{PlayerAccountIDs: "foo", LevelUUID: "level1", Value: 10, ValueType: 0}},
		},
	}
	assert.Equal(t, expectedLeaderboard, leaderboards)
}
