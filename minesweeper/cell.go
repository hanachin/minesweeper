package minesweeper

import (
	"errors"
	"fmt"
	"./term"
)

// FIXME
var debug bool = false

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
		fmt.Print("X")
	} else {
		colorMap := []term.Color{ term.ColorWhite,
			term.ColorCyan, term.ColorBlue, term.ColorGreen,
			term.ColorYellow, term.ColorMagenta, term.ColorWhite,
			term.ColorWhite, term.ColorWhite}
		bombCount := c.BombCount()
		term.SetForegroundColor(colorMap[bombCount])
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
		term.SetForegroundColor(term.ColorRed)
		fmt.Print("X")
	} else if c.BombCount() == 0 {
		fmt.Print(" ")
	} else {
		colorMap := []term.Color{ term.ColorBlack,
			term.ColorCyan, term.ColorBlue, term.ColorGreen,
			term.ColorYellow, term.ColorMagenta, term.ColorBlack,
			term.ColorBlack, term.ColorBlack}
		bombCount := c.BombCount()
		term.SetForegroundColor(colorMap[bombCount])
		fmt.Printf("%d", bombCount)
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
