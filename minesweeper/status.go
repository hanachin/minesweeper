package minesweeper

import (
	"fmt"
	"strconv"
	"./term"
)

type Status struct {
	X int
	Y int
	Width int
}

func (s *Status) ShowMessage(m string, color term.Color) {
	term.ResetColor()
	term.SetForegroundColor(color)
	term.SetCursor(s.X + 1, s.Y + 2)
	fmt.Printf("%-" + strconv.Itoa(s.Width) + "s\n", m)
}
