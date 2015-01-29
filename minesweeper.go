package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
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

func (m *Map) StartPoint() *Cell {
	return m.Cells[0]
}

type Status struct {
	X int
	Y int
	Width int
}

func (s *Status) ShowMessage(m string) {
	term.ResetColor()
	term.SetCursor(s.X + 1, s.Y + 1)
	fmt.Printf("%-" + strconv.Itoa(s.Width) + "s\n", m)
}

type Game struct {
	CurrentCell *Cell
	Map *Map
	Status *Status
}

func NewGame() *Game {
	g := new(Game)
	g.Status = new(Status)
	return g
}

func (g *Game) SetSquareMap(cols, rows, bomb int) {
	m := NewSquareMap(cols, rows)
	m.PutBomb(bomb)
	g.Map = m
	g.Status.Y = rows
	g.Status.Width = cols
}

func (g *Game) Move(dir int) {
	if (g.CurrentCell.Neighbors[dir] != nil) {
		g.CurrentCell = g.CurrentCell.Neighbors[dir]
	}
}

func (g *Game) Start() {
	term.WithGameMode(func () {
		g.CurrentCell = g.Map.StartPoint()
		g.Map.Show()
	Loop:
		for {
			term.SetCursor(g.CurrentCell.X + 1, g.CurrentCell.Y + 1)

			switch term.Getc() {
			case 'q':
				g.Status.ShowMessage("bye bye!")
				break Loop
			case 'w':
				g.Move(DirUp)
			case 'a':
				g.Move(DirLeft)
			case 's':
				g.Move(DirBottm)
			case 'd':
				g.Move(DirRight)
			case ' ':
				err := g.CurrentCell.Open()
				if err != nil {
					term.SetCursor(g.Status.X + 1, g.Status.Y + 1)
					term.ResetColor()
					term.SetForegroundColor(term.ColorRed)
					fmt.Println(err)
					break Loop
				}
			case 'f':
				if !g.CurrentCell.IsOpened {
					g.CurrentCell.ToggleDangerSign()
				}
			}

		}
	})
}

func main() {
	g := NewGame()
	g.SetSquareMap(40, 20, 100)
	g.Start()
}
