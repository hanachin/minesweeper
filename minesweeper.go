package main

import (
	"fmt"
	"math/rand"
	"time"
	"./term"
)

type Cell struct {
	IsBomb bool
	IsOpened bool
	X int
	Y int
	Neighbors []*Cell
}

func (c1 *Cell) Connect(c2 *Cell) {
	c1.Neighbors = append(c1.Neighbors, c2)
}

func (c *Cell) BombCount() int {
	count := 0
	for i := 0; i < len(c.Neighbors); i++ {
		if (c.Neighbors[i].IsBomb) {
			count += 1
		}
	}
	return count
}

func (c *Cell) Show() {
	if c.IsOpened {
		term.SetBackgroundColor(term.ColorWhite)
	} else {
		term.SetBackgroundColor(term.ColorBlack)
	}
	term.SetForegroundColor(term.ColorRed)
	term.SetCursor(c.X + 1, c.Y + 1)
	if c.IsBomb {
		term.Print("f")
	} else {
		fmt.Printf("%d", c.BombCount())
	}

}

type Map struct {
	Cells []*Cell
}

func NewSquareMap(cols, rows int) *Map {
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
			// up
			if (y - 1 >= 0) {
				c.Connect(&cells[y - 1][x])
			}
			// up right
			if (y - 1 >= 0 && x + 1 < cols) {
				c.Connect(&cells[y - 1][x + 1])
			}
			// right
			if (x + 1 < cols) {
				c.Connect(&cells[y][x + 1])
			}
			// down right
			if (x + 1 < cols && y + 1 < rows) {
				c.Connect(&cells[y + 1][x + 1])
			}
			// down
			if (y + 1 < rows) {
				c.Connect(&cells[y + 1][x])
			}
			// down left
			if (y + 1 < rows && x - 1 >= 0) {
				c.Connect(&cells[y + 1][x - 1])
			}
			// left
			if (x - 1 >= 0) {
				c.Connect(&cells[y][x - 1])
			}
			// left up
			if (x - 1 >= 0 && y - 1 >= 0) {
				c.Connect(&cells[y - 1][x - 1])
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

type CursorPoint struct {
	X int
	Y int
}

func main() {
	term.WithGameMode(func () {
		cursor := new(CursorPoint)
		cols := 4
		rows := 7
		m := NewSquareMap(cols, rows)
		m.PutBomb(5)
		m.Show()
	Loop:
		for {
			term.SetCursor(cursor.X + 1, cursor.Y + 1)
			c := term.Getc()

			switch c {
			case 'q':
				break Loop
			case 'w':
				if (cursor.Y - 1 >= 0) {
					cursor.Y -= 1
				}
			case 'a':
				if (cursor.X - 1 >= 0) {
					cursor.X -= 1
				}
			case 's':
				if (cursor.Y + 1 < rows) {
					cursor.Y += 1
				}
			case 'd':
				if (cursor.X + 1 < cols) {
					cursor.X += 1
				}
			}
		}
		term.SetCursor(1, rows + 1)
		term.ResetColor()
	})
}
