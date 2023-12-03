package main

import (
	"bytes"
	"os"
)

type Grid [][]byte

func LoadFile(filename string) (g Grid) {
	input, err := os.ReadFile(filename)
	PanicOnError(err)
	return load(input)
}

func LoadString(inputStr string) (g Grid) {
	return load([]byte(inputStr))
}

func load(input []byte) (g Grid) {
	return bytes.Split(bytes.TrimSpace(input), []byte("\n"))
}

func (g Grid) String() string {
	text := make([][]byte, len(g))
	for i, row := range g {
		text[i] = make([]byte, len(row)+2)
		text[i][0] = '|'
		copy(text[i][1:], row)
		text[i][len(row)+1] = '|'
	}
	return string(bytes.Join(text, []byte("\n")))
}

func (g Grid) Width() int {
	// Assume all the rows are the same length
	return len(g[0])
}

func (g Grid) Height() int {
	return len(g)
}

func (g Grid) Clamp(x, y int) (int, int) {
	if x < 0 {
		x = 0
	}
	if x > g.Width() {
		x = g.Width()
	}
	if y < 0 {
		y = 0
	}
	if y > g.Height() {
		y = g.Height()
	}
	return x, y
}

func (g Grid) SubGrid(x, y, w, h int) Grid {
	//log.Printf("SubGrid: x=%d y=%d w=%d h=%d", x, y, w, h)

	if x < 0 {
		w += x
		x = 0
	}
	if x > g.Width() {
		x = 0
		w = 0
	}
	if y < 0 {
		h += y
		y = 0
	}
	if y > g.Height() {
		y = 0
		h = 0
	}

	if x+w > g.Width() {
		w = g.Width() - x
	}
	w = max(w, 0) // no negative width (should never happen)

	if y+h > g.Height() {
		h = g.Height() - y
	}
	h = max(h, 0) // no negative height (should never happen)

	//log.Printf("SubGrid: x=%d y=%d w=%d h=%d", x, y, w, h)

	rows := make([][]byte, h)
	for i := 0; i < h; i++ {
		rows[i] = g[y+i][x : x+w]
	}
	return rows
}

func (g Grid) ContainsSymbol() (symbol byte) {
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			switch g[y][x] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				// this is a number
				continue
			case '.':
				// this is a space - it doesn't count as a symbol
				continue
			default:
				return g[y][x]
			}
		}
	}
	return ' ' // using a space to represent "not found"
}

type Point struct {
	X int
	Y int
}

func (g Grid) Find(symbol byte) []Point {
	found := make([]Point, 0)
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if g[y][x] == symbol {
				found = append(found, Point{x, y})
			}
		}
	}
	return found
}
