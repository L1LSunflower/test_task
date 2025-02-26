package symbols

import (
	"embed"
	"fmt"
	"strconv"
	"sync"

	"github.com/releaseband/golang-developer-test/internal/configs/reader"
)

//go:embed symbols.txt
var symbols embed.FS

const skipSymbol = -1

func parseReels(data [][]string) ([]Symbols, error) {
	symbolsLength := 0
	if len(data) > 0 {
		symbolsLength = len(data[0])
	}
	parsedData := make([]Symbols, symbolsLength)

	wg := new(sync.WaitGroup)
	var err error
	for col := 0; col < symbolsLength; col++ {
		wg.Add(1)

		go func(col int) {
			defer wg.Done()

			for _, row := range data {
				var num int
				num, err = strconv.Atoi(row[col])
				if err != nil {
					err = fmt.Errorf("strconv.Atoi(row[col]): %w", err)
				}

				if num == skipSymbol {
					continue
				}

				parsedData[col] = append(parsedData[col], num)
			}

		}(col)
	}

	wg.Wait()

	return parsedData, err
}

// ReadReels - read symbols from file
func ReadReels() ([]Symbols, error) {
	// обрати внимание, что в файле symbols.txt символы разделены через \t
	// и что в конце каждой строки есть \n
	// символ -1 нужен только для выравнивания таблицы
	data, err := reader.Read(symbols, "symbols.txt")
	if err != nil {
		return nil, fmt.Errorf("reader.Read(): %w", err)
	}

	symbols, err := parseReels(data)
	if err != nil {
		return nil, fmt.Errorf("parseReels(): %w", err)
	}

	return symbols, nil
}
