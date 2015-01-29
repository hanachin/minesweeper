package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"./term"
)

var debug bool = false

const (
	DirUp = iota
	DirUpperRight
	DirRight
	DirLowerRight
	DirBottm
	DirLowerLeft
	DirLeft
	DirUpperLeft
)

type Cell struct {
	IsBomb bool
	IsOpened bool
	DangerSign bool
	X int
	Y int
	Neighbors map[int]*Cell
}

func (c1 *Cell) Connect(dir int, c2 *Cell) {
	c1.Neighbors[dir] = c2
}

func (c *Cell) BombCount() int {
	count := 0
	for _, nc := range c.Neighbors {
		if (nc.IsBomb) {
			count += 1
		}
	}
	return count
}

func (c *Cell) DebugShow() {
	term.SetCursor(c.X + 1, c.Y + 1)

	if c.IsOpened {
		term.SetForegroundColor(term.ColorBlack)
		term.SetBackgroundColor(term.ColorWhite)
	} else {
		term.SetForegroundColor(term.ColorWhite)
		term.SetBackgroundColor(term.ColorBlack)
	}

	if c.DangerSign {
		term.SetForegroundColor(term.ColorRed)
		fmt.Print("f")
	} else if c.IsBomb {
		term.SetForegroundColor(term.ColorCyan)
		fmt.Print("b")
	} else {
		fmt.Printf("%d", c.BombCount())
	}
}

func (c *Cell) Show() {
	if debug {
		c.DebugShow()
		return
	}

	term.SetCursor(c.X + 1, c.Y + 1)

	if c.IsOpened {
		term.SetForegroundColor(term.ColorBlack)
		term.SetBackgroundColor(term.ColorWhite)
	} else {
		term.SetForegroundColor(term.ColorWhite)
		term.SetBackgroundColor(term.ColorBlack)
	}

	if !c.IsOpened {
		if c.DangerSign {
			term.SetForegroundColor(term.ColorRed)
			fmt.Print("f")
		} else {
			fmt.Print(" ")
		}
		return
	}

	if c.DangerSign {
		term.SetForegroundColor(term.ColorRed)
		fmt.Print("f")
	} else if c.IsBomb {
		term.SetForegroundColor(term.ColorCyan)
		fmt.Print("b")
	} else if c.BombCount() == 0 {
		fmt.Print(" ")
	} else {
		fmt.Printf("%d", c.BombCount())
	}

}

func (c *Cell) Open() error {
	if c.IsOpened {
		return nil
	}

	c.IsOpened = true
	c.Show()

	if c.IsBomb {
		return errors.New("bomb! you dead.")
	}

	if c.BombCount() == 0 {
		for _, nc := range c.Neighbors {
			nc.Open()
		}
	}
	return nil
}

func (c *Cell) ToggleDangerSign() {
	c.DangerSign = !c.DangerSign
	c.Show()
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

func main() {
	term.WithGameMode(func () {
		cols := 4
		rows := 7
		m := NewSquareMap(cols, rows)
		m.PutBomb(5)
		m.Show()

		currentCell := m.Cells[0]
	Loop:
		for {
			term.SetCursor(currentCell.X + 1, currentCell.Y + 1)
			c := term.Getc()

			switch c {
			case 'q':
				break Loop
			case 'w':
				if (currentCell.Neighbors[DirUp] != nil) {
					currentCell = currentCell.Neighbors[DirUp]
				}
			case 'a':
				if (currentCell.Neighbors[DirLeft] != nil) {
					currentCell = currentCell.Neighbors[DirLeft]
				}
			case 's':
				if (currentCell.Neighbors[DirBottm] != nil) {
					currentCell = currentCell.Neighbors[DirBottm]
				}
			case 'd':
				if (currentCell.Neighbors[DirRight] != nil) {
					currentCell = currentCell.Neighbors[DirRight]
				}

			case ' ':
				err := currentCell.Open()
				if err != nil {
					term.SetCursor(1, rows + 1)
					term.ResetColor()
					term.SetForegroundColor(term.ColorRed)
					fmt.Println(err)
					break Loop
				}
			case 'f':
				currentCell.ToggleDangerSign()
			}

		}
		term.SetCursor(1, rows + 2)
		term.ResetColor()
	})
}
