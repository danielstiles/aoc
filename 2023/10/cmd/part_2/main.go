package main

import (
	"bufio"
	"os"

	"log/slog"

	"github.com/danielstiles/aoc/2023/10/internal/pipes"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		slog.Error("Could not read file", slog.Any("error", err))
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	area := &pipes.Area{}
	for fileScanner.Scan() {
		area.AddRow([]byte(fileScanner.Text()))
	}
	area.CalcPath()
	slog.Info("Answer", slog.Int("interior", area.GetInteriorCount()))
}
