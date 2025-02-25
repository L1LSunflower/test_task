package calculator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/releaseband/golang-developer-test/internal/configs/lines"
	"github.com/releaseband/golang-developer-test/internal/configs/paytable"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/game/win"
)

func TestCalculator_Calculate_V3_RealPayTable(t *testing.T) {
	testLines, err := lines.ReadLines()
	if err != nil {
		t.Errorf("failed to read lines with error: %s", err)
	}

	realPayTable, err := paytable.ReadPayTable()
	if err != nil {
		panic(fmt.Sprintf("Failed to load real paytable: %v", err))
	}

	tests := []struct {
		name          string
		spinSymbols   symbols.Reels
		expectedWins  []win.Win
		expectedError bool
	}{
		{
			name: "Invalid spinSymbols - less than 3 rows",
			spinSymbols: symbols.Reels{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
			},
			expectedWins:  nil,
			expectedError: true,
		},
		{
			name: "No wins on any lines",
			spinSymbols: symbols.Reels{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
			},
			expectedWins:  []win.Win{},
			expectedError: false,
		},
		{
			name: "Win on line 1 - 3 symbols of type 1 (based on pay_table.txt)",
			spinSymbols: symbols.Reels{
				{1, 1, 1, 4, 5},
				{6, 7, 8, 9, 10},
				{1, 2, 3, 4, 5},
			},
			expectedWins: []win.Win{
				win.NewWin(50, []int{1, 1, 1}, 1),
			},
			expectedError: false,
		},
		{
			name: "Win on line 2 - 4 symbols of type 2 (based on pay_table.txt)",
			spinSymbols: symbols.Reels{
				{1, 2, 3, 4, 5},
				{2, 2, 2, 2, 10},
				{1, 2, 3, 4, 5},
			},
			expectedWins: []win.Win{
				win.NewWin(200, []int{2, 2, 2, 2}, 2),
			},
			expectedError: false,
		},
		{
			name: "Win on line 3 - 5 symbols of type 3 (based on pay_table.txt)",
			spinSymbols: symbols.Reels{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{3, 3, 3, 3, 3},
			},
			expectedWins: []win.Win{
				win.NewWin(200, []int{3, 3, 3, 3, 3}, 3),
			},
			expectedError: false,
		},
		{
			name: "Win on line 4 - Diagonal line win - 3 symbols of type 7 (based on pay_table.txt)",
			spinSymbols: symbols.Reels{
				{7, 2, 3, 4, 1},
				{6, 7, 8, 9, 10},
				{1, 2, 7, 4, 5},
			},
			expectedWins: []win.Win{
				win.NewWin(10, []int{7, 7, 7}, 7),
			},
			expectedError: false,
		},
		{
			name: "Win on line 5 - Another diagonal line win - 3 symbols of type 6 (based on pay_table.txt)",
			spinSymbols: symbols.Reels{
				{1, 6, 7, 4, 1},
				{6, 7, 6, 9, 10},
				{7, 2, 3, 4, 6},
			},
			expectedWins: []win.Win{
				win.NewWin(10, []int{7, 7, 7}, 7),
			},
			expectedError: false,
		},
		{
			name: "Wins on multiple lines - lines 1 and 2 (based on pay_table.txt)",
			spinSymbols: symbols.Reels{
				{1, 1, 1, 4, 5},
				{2, 2, 2, 2, 10},
				{1, 2, 3, 4, 5},
			},
			expectedWins: []win.Win{
				win.NewWin(50, []int{1, 1, 1}, 1),
				win.NewWin(200, []int{2, 2, 2, 2}, 2),
			},
			expectedError: false,
		},
		{
			name: "Win on line 1 with WILD - 3 symbols of type 1 (1, 1, 0) (based on pay_table.txt)",
			spinSymbols: symbols.Reels{
				{1, 1, 0, 4, 5},
				{6, 7, 8, 9, 10},
				{1, 2, 3, 4, 5},
			},
			expectedWins: []win.Win{
				win.NewWin(50, []int{1, 1, 0}, 1),
			},
			expectedError: false,
		},
		{
			name: "No win - WILD not helping to form a win",
			spinSymbols: symbols.Reels{
				{5, 1, 0, 3, 4},
				{6, 7, 8, 9, 10},
				{1, 2, 3, 4, 5},
			},
			expectedWins:  []win.Win{},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := NewCalculator(testLines, realPayTable)
			wins, err := calc.Calculate(tt.spinSymbols)

			if tt.expectedError {
				require.Error(t, err, "Expected error but got nil")
			} else {
				require.NoError(t, err, "Unexpected error")
			}

			require.ElementsMatch(t, tt.expectedWins, wins, "Wins mismatch")
		})
	}
}
