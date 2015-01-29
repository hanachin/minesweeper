package minesweeper

import (
	"math/rand"
	"time"
	"./term"
)

type Map struct {
	Cells []*Cell
}

func NewRectangleMap(cols, rows int) *Map {
	m := new(Map)
	m.Cells = make([]*Cell, cols * rows)
	cells := make([][]Cell, rows)
	for y := 0; y < rows; y++ {
		cells[y] = make([]Cell, cols)
	}
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			c := &cells[y][x]
			c.X = x
			c.Y = y
			c.Neighbors = map[int]*Cell{}
			// up
			if (y - 1 >= 0) {
				c.Connect(DirUp, &cells[y - 1][x])
			}
			// up right
			if (y - 1 >= 0 && x + 1 < cols) {
				c.Connect(DirUpperRight, &cells[y - 1][x + 1])
			}
			// right
			if (x + 1 < cols) {
				c.Connect(DirRight, &cells[y][x + 1])
			}
			// down right
			if (x + 1 < cols && y + 1 < rows) {
				c.Connect(DirLowerRight, &cells[y + 1][x + 1])
			}
			// down
			if (y + 1 < rows) {
				c.Connect(DirBottm, &cells[y + 1][x])
			}
			// down left
			if (y + 1 < rows && x - 1 >= 0) {
				c.Connect(DirLowerLeft, &cells[y + 1][x - 1])
			}
			// left
			if (x - 1 >= 0) {
				c.Connect(DirLeft, &cells[y][x - 1])
			}
			// left up
			if (x - 1 >= 0 && y - 1 >= 0) {
				c.Connect(DirUpperLeft, &cells[y - 1][x - 1])
			}
			m.Cells[y * cols + x] = c
		}
	}
	return m
}

func (m *Map) PutBomb(n int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rs := map[int]bool{}
	max := len(m.Cells)
	for len(rs) < n {
		bombPos := r.Intn(max)
		rs[bombPos] = true
		m.Cells[bombPos].IsBomb = true
	}
}

func (m *Map) Show() {
	term.Clear()
	for _, c := range m.Cells {
		c.Show()
	}
}

func (m *Map) StartPoint() *Cell {
	return m.Cells[0]
}

func (m *Map) CountOpenedCells() int {
	openedCells := 0
	for _, c := range m.Cells {
		if c.IsOpened {
			openedCells += 1
		}
	}
	return openedCells
}

func (m *Map) CountCells() int {
	return len(m.Cells)
}
