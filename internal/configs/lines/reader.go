package lines

import (
	"embed"
	"fmt"
	"strconv"

	"github.com/releaseband/golang-developer-test/internal/configs/reader"
)

//go:embed lines.txt
var lines embed.FS

func parseLine(data []string) (*Line, error) {
	lineInts := make([]int, len(data))

	for index, numStr := range data {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
		lineInts[index] = num
	}

	return NewLine(lineInts), nil
}

func ReadLines() (Lines, error) {
	data, err := reader.Read(lines, "lines.txt")
	if err != nil {
		return nil, fmt.Errorf("reader.Read(): %w", err)
	}

	resp := make([]Line, len(data))
	for i, str := range data {
		line, err := parseLine(str)
		if err != nil {
			return nil, fmt.Errorf("parseLines(): %w", err)
		}

		resp[i] = *line
	}

	return resp, nil
}
