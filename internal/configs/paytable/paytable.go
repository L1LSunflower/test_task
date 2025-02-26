package paytable

import (
	"fmt"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
)

// Payout - таблица выплаты определенного символа
type Payout []uint64

// PayTable - таблица выплат всех символов
type PayTable struct {
	symbolPayouts map[symbols.Symbol]Payout
}

func NewPayTable(symbolPayouts map[symbols.Symbol]Payout) *PayTable {
	return &PayTable{symbolPayouts: symbolPayouts}
}

func (p *PayTable) Get(s symbols.Symbol, index int) (uint64, error) {
	payRow, ok := p.symbolPayouts[s]
	if !ok {
		return 0, fmt.Errorf("p.symbolPayouts[s]: this symbol '%d' does not exist", s)
	}

	if len(payRow) < index {
		return 0, fmt.Errorf("does not exist index: %d", index)
	}

	return payRow[index], nil
}
