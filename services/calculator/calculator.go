package calculator

import (
	"fmt"

	"github.com/releaseband/golang-developer-test/internal/configs/lines"
	"github.com/releaseband/golang-developer-test/internal/configs/paytable"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/game/win"
)

// WILD - специальный символ, который может заменить любой другой символ
// он не имеет своего выигрыша, но может увеличить выигрыш за счет замены другого символа
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
	if len(spinSymbols) < 3 {
		return nil, fmt.Errorf("spin symbols length less than 3, spinSymbols length: %d", len(spinSymbols))
	}

	var wins []win.Win

	for lineNum := 0; lineNum < len(c.lines); lineNum++ {
		var (
			streak        = 0
			streakSymbols []int
		)

		rowToCheck := c.getRow(spinSymbols, lineNum)
		streakSymbols, streak = c.checkRow(rowToCheck)

		if streak == skip {
			continue
		}

		amount, err := c.payTable.Get(streakSymbols[0], streak)
		if err != nil {
			return wins, err
		}

		wins = append(wins, win.NewWin(amount, streakSymbols, streakSymbols[0]))
	}

	return wins, nil
}

func (c *Calculator) getRow(symbols symbols.Reels, lineNum int) []int {
	indices := c.lines[lineNum].GetIndices()
	finalRow := make([]int, len(indices))

	for col, row := range indices {
		finalRow[col] = symbols[row][col]
	}

	return finalRow
}

func (c *Calculator) checkRow(row symbols.Symbols) ([]int, int) {
	if len(row) == 0 {
		return nil, skip
	}

	count := 1
	pVal := row[0]

	if pVal == WILD {
		return nil, skip
	}

	for col := 1; col < len(row); col++ {
		if pVal != row[col] && col == 1 && row[col] != WILD {
			return nil, skip
		}

		if (pVal == row[col] || row[col] == WILD) && count == col {
			count++
		}

	}

	return row[:count], count - 1
}
