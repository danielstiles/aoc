package main

import (
	"bufio"
	"log/slog"
	"os"
	"time"

	"github.com/danielstiles/aoc/2024/20/internal/parts"
)

func main() {
	start := time.Now()
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		slog.Error("Could not read file", slog.Any("error", err))
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readTime := time.Now()
	total := parts.Process2(lines)
	processTime := time.Now()
	slog.Info(
		"Answer",
		slog.Duration("read time", readTime.Sub(start)),
		slog.Duration("process time", processTime.Sub(readTime)),
		slog.Duration("total time", processTime.Sub(start)),
		slog.Int("total", total),
	)
}
