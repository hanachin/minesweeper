package minesweeper

import (
	"fmt"
	"strconv"
	"./term"
)

const helpMessage string = "help: wasd(move) space(open) f(flag)"

type Status struct {
	X int
	Y int
	Width int
}

func (s *Status) ShowMessage(m string, color term.Color) {
	term.ResetColor()
	term.SetForegroundColor(color)
	term.SetCursor(s.X + 1, s.Y + 2)
	w := s.Width
	if w < len(helpMessage) {
		w = len(helpMessage)
	}
	fmt.Printf("%-" + strconv.Itoa(w) + "s\n", m)
}

func (s *Status) ShowHelp() {
	s.ShowMessage(helpMessage, term.ColorBlack)
}

func (s *Status) ShowDangerSignCount(count, total int) {
	term.ResetColor()
	term.SetForegroundColor(term.ColorBlack)
	term.SetCursor(s.X + 1, s.Y + 1)
	mes := fmt.Sprintf("bomb: %d / %d", total - count, total)
	fmt.Printf("%-" + strconv.Itoa(s.Width) + "s\n", mes)
}
