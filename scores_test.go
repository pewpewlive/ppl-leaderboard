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

func TestGetScoreMapFromScores(t *testing.T) {
	scores := []ppl_types.HofEntry{}
	accounts := []ppl_types.AccountInfo{}
	leaderboard := GetLeaderboardsFromScores(scores, accounts)
	expectedLeaderboard := []LevelLeaderboard{}
	assert.Equal(t, expectedLeaderboard, leaderboard)
}
