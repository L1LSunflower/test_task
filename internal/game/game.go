package game

import (
	"fmt"

	"github.com/releaseband/golang-developer-test/internal/game/result"
	"github.com/releaseband/golang-developer-test/internal/rng"
)

// RoundCost - функция, которая возвращет стоимость одного раунда
func RoundCost(linesCount int) uint64 {
	return uint64(linesCount)
}

type Slot struct {
	generator  Generator
	calculator Calculator
	roundCost  uint64
}

func newSlot(generator Generator, calculator Calculator, roundCost uint64) *Slot {
	return &Slot{generator: generator, calculator: calculator, roundCost: roundCost}
}

func (s *Slot) Spin(rng rng.RNG) (*result.Round, error) {
	gameField, err := s.generator.Generate(rng)
	if err != nil {
		return nil, fmt.Errorf("failed to generate game field: %w", err)
	}

	calculationResult, err := s.calculator.Calculate(gameField)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate wins: %w", err)
	}

	var totalWinAmount uint64 = 0
	for _, calcRes := range calculationResult {
		totalWinAmount += calcRes.Amount()
	}

	return result.NewRound(gameField, calculationResult, totalWinAmount), nil
}

func (s *Slot) RoundCost() uint64 {
	return s.roundCost
}
