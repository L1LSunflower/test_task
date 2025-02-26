package calculator

import (
	"github.com/releaseband/golang-developer-test/internal/configs/lines"
	"github.com/releaseband/golang-developer-test/internal/configs/paytable"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/game/win"
)

const WILD = symbols.Symbol(0)
const skip = -1

type Calculator struct {
	lines    lines.Lines
	payTable *paytable.PayTable
}

func NewCalculator(lines lines.Lines, payTable *paytable.PayTable) *Calculator {
	return &Calculator{lines: lines, payTable: payTable}
}

func (c *Calculator) Calculate(spinSymbols symbols.Reels) ([]win.Win, error) {
	var wins []win.Win

	for lineNum := 0; lineNum < len(c.lines); lineNum++ {
		rowToCheck := c.getRow(spinSymbols, lineNum)
		streakSymbols, streak := c.checkRow(rowToCheck)

		amount, err := c.payTable.Get(streakSymbols[0], streak-1)
		if err != nil {
			return wins, err
		}

		if amount == 0 {
			continue
		}

		wins = append(wins, win.NewWin(amount, streakSymbols, streakSymbols[0]))
	}

	return wins, nil
}

func (c *Calculator) getRow(symbols symbols.Reels, lineNum int) []int {
	indices := c.lines[lineNum].GetIndices()
	finalRow := make([]int, len(indices))

	for col, rowIndex := range indices {
		finalRow[col] = symbols[col][rowIndex]
	}

	return finalRow
}

func (c *Calculator) checkRow(row symbols.Symbols) ([]int, int) {
	if len(row) == 0 {
		return nil, skip
	}

	count := 0
	pVal := -1

	// Count all leading WILDs
	for col := 0; col < len(row); col++ {
		if row[col] == WILD {
			count++
		} else {
			pVal = row[col]
			break
		}
	}

	if pVal == -1 {
		return row[:count], count
	}

	// Count consecutive matching symbols or WILDs
	for col := count; col < len(row); col++ {
		if row[col] == pVal || row[col] == WILD {
			count++
		} else {
			break
		}
	}

	return row[:count], count
}
