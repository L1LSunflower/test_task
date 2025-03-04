package game

import (
	"github.com/releaseband/golang-developer-test/internal/configs/lines"
	"github.com/releaseband/golang-developer-test/internal/configs/paytable"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/services/calculator"
	"github.com/releaseband/golang-developer-test/services/generator"
)

const rowsCount = 3

func New() (*Slot, error) {
	// read lines configs
	cfgLines, err := lines.ReadLines()
	if err != nil {
		return nil, err
	}

	// read pay table configs
	payTable, err := paytable.ReadPayTable()
	if err != nil {
		return nil, err
	}

	// read game tapes
	gameTapes, err := symbols.ReadReels()
	if err != nil {
		return nil, err
	}

	g := generator.NewSymbols(rowsCount, gameTapes)
	c := calculator.NewCalculator(cfgLines, payTable)

	// create services
	return newSlot(g, c, uint64(len(cfgLines))), nil
}
