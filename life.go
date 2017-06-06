package main

import (
	"crypto/rand"
	"encoding/binary"
	"bytes"
	"math"
	"fmt"
	"time"
	"os"
	"os/signal"
	"github.com/fatih/color"
)

type Point struct {
	X, Y int
}

func (p Point) Neighbours() []Point {
	return []Point{
		{p.X - 1, p.Y - 1},
		{p.X, p.Y - 1},
		{p.X + 1, p.Y - 1},
		{p.X - 1, p.Y},
		{p.X + 1, p.Y},
		{p.X - 1, p.Y + 1},
		{p.X, p.Y + 1},
		{p.X + 1, p.Y + 1},
	}
}

type GameBoard struct {
	cells      map[Point]bool
	MinX, MaxX int
	MinY, MaxY int
}

func (b *GameBoard) Refresh() {
	for p := range b.cells {
		b.MinX = min(b.MinX, p.X)
		b.MaxX = max(b.MaxX, p.X)
		b.MinY = min(b.MinY, p.Y)
		b.MaxY = max(b.MaxY, p.Y)
	}
}

func (b *GameBoard) String() string {
	var buffer bytes.Buffer

	for y := b.MinY; y <= b.MaxY; y++ {
		for x := b.MinX; x <= b.MaxX; x ++ {
			if v, ok := b.cells[Point{x, y}]; !ok || !v {
				buffer.WriteString(" ")
			} else {
				buffer.WriteString("•")
			}
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}

func (board *GameBoard) Height() int {
	return int(math.Abs(float64(board.MaxY - board.MinY)))
}

func (board *GameBoard) Width() int {
	return int(math.Abs(float64(board.MaxX - board.MinX)))
}

type compareFunc func(int, int) bool

var (
	less    = func(a, b int) bool { return a < b }
	greater = func(a, b int) bool { return a > b }
)

func min(a, b int) int {
	return compare(less, a, b)
}

func max(a, b int) int {
	return compare(greater, a, b)
}

func compare(f compareFunc, a, b int) int {
	if f(a, b) {
		return a
	}
	return b
}

func newBoardWithCells(cells map[Point]bool) *GameBoard {
	b := &GameBoard{cells, math.MaxInt64, math.MinInt64, math.MaxInt64, math.MinInt64}
	b.Refresh()
	return b
}

func New(width int, height int) *GameBoard {
	board := make(map[Point]bool)
	var num int8

	for y := 0; y < height; y++ {
		for x := 0; x < width; x ++ {
			binary.Read(rand.Reader, binary.LittleEndian, &num)
			if num&1 == 1 {
				board[Point{x, y}] = true
			}
		}
	}

	return newBoardWithCells(board)
}
func Next(b *GameBoard) *GameBoard {
	neighbours := make(map[Point]int)

	for p := range b.cells {
		for _, n := range p.Neighbours() {
			if v, ok := neighbours[n]; ok {
				neighbours[n] = v + 1
			} else {
				neighbours[n] = 1
			}
		}
	}

	cells := make(map[Point]bool)
	for p, v := range neighbours {
		switch v {
		case 2:
			if _, ok := b.cells[p]; ok {
				cells[p] = true
			}
		case 3:
			cells[p] = true
		}
	}

	return newBoardWithCells(cells)
}

func NextOld(b *GameBoard) *GameBoard {
	next := make(map[Point]bool)

	for y := b.MinY - 1; y <= b.MaxY+1; y++ {
		for x := b.MinX - 1; x <= b.MaxX+1; x ++ {
			p := Point{x, y}
			neighbours := countAliveNeighbours(b, p)
			switch {
			case neighbours == 2:
				if _, ok := b.cells[p]; ok {
					next[p] = true
				}
			case neighbours == 3:
				next[p] = true
			}
		}
	}

	return newBoardWithCells(next)
}

func countAliveNeighbours(b *GameBoard, p Point) int {
	var aliveNeighbours int

	for _, n := range p.Neighbours() {
		if v, ok := b.cells[n]; ok && v {
			aliveNeighbours++
		}
	}

	return aliveNeighbours
}

func main() {
	go func() {
		board := New(40, 20)
		for step := 0; ; step++ {
			PrintFrame(step, board)
			time.Sleep(500 * time.Millisecond)
			board = Next(board)
		}
	}()

	killSig := make(chan os.Signal)
	signal.Notify(killSig, os.Interrupt, os.Kill)
	_ = <-killSig
}

func PrintFrame(step int, b *GameBoard) {
	fmt.Print("\033[2J")
	fmt.Print("\033[;H")
	color.Green("Step: %d \t width: %d | height: %d", step, b.Width(), b.Height())
	color.Red("%v", b)
}
