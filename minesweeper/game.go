package minesweeper

import (
	"./term"
)

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
				g.Status.ShowMessage("bye bye!", term.ColorBlack)
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
					g.Status.ShowMessage(err.Error(), term.ColorRed)
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
