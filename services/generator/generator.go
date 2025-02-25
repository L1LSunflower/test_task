package generator

import (
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/rng"
)

type Symbols struct {
	rowsCount int
	gameTapes []symbols.Symbols
}

func NewSymbols(rowsCount int, gameTapes []symbols.Symbols) *Symbols {
	return &Symbols{rowsCount: rowsCount, gameTapes: gameTapes}
}

func (s *Symbols) Generate(rng rng.RNG) (symbols.Reels, error) {
	const numberOfReels = 5

	gameField := make(symbols.Reels, s.rowsCount)
	for i := range gameField {
		gameField[i] = make([]int, numberOfReels)
	}

	for colI := 0; colI < numberOfReels; colI++ {
		symbolTape := s.gameTapes[colI]

		startIndex := rng.Random(0, uint32(len(symbolTape)))
		reelColumnSymbols := s.GetReelSymbols(colI, int(startIndex))

		for rowI := 0; rowI < s.rowsCount; rowI++ {
			gameField[rowI][colI] = reelColumnSymbols[rowI]
		}
	}

	return gameField, nil
}

func (s *Symbols) GetReelSymbols(reelIndex int, rowIndex int) symbols.Symbols {
	symbolTape := s.gameTapes[reelIndex]
	tapeLength := len(symbolTape)
	reelSymbols := make(symbols.Symbols, 3)

	for i := 0; i < s.rowsCount; i++ {
		indexInTape := (rowIndex + i) % tapeLength
		symbol := symbolTape[indexInTape]
		reelSymbols[i] = symbol
	}

	return reelSymbols
}
