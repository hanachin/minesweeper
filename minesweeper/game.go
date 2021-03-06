package minesweeper

import (
	"./term"
)

type Game struct {
	CurrentCell *Cell
	Map *Map
	Status *Status
	Bomb int
	DangerSignCount int
}

func NewGame() *Game {
	g := new(Game)
	g.Status = new(Status)
	return g
}

func (g *Game) SetRectangleMap(cols, rows, bomb int) {
	m := NewRectangleMap(cols, rows)
	m.PutBomb(bomb)
	g.Bomb = bomb
	g.Map = m
	g.Status.Y = rows
	g.Status.Width = cols
}

func (g *Game) SetTemplateMap(template []string, bomb int) {
	m := NewTemplateMap(template)
	m.PutBomb(bomb)
	g.Bomb = bomb
	g.Map = m
	g.Status.Y = len(template)
	g.Status.Width = len(template[0])
}


func (g *Game) Move(dir int) {
	if (g.CurrentCell.Neighbors[dir] != nil) {
		g.CurrentCell = g.CurrentCell.Neighbors[dir]
	}
}

func (g *Game) ShowMessage(m string, color term.Color) {
	g.Status.ShowMessage(m, color)
}

func (g *Game) ShowHelp() {
	g.Status.ShowHelp()
}

func (g *Game) ShowDangerSignCount() {
	g.Status.ShowDangerSignCount(g.DangerSignCount, g.Bomb)
}

func (g *Game) IsClear() bool {
	return g.Bomb + g.Map.CountOpenedCells() == g.Map.CountCells()
}

func (g *Game) Start() {
	term.WithGameMode(func () {
		g.CurrentCell = g.Map.StartPoint()
		g.Map.Show()
		g.ShowHelp()
	Loop:
		for {
			g.ShowDangerSignCount()
			term.SetCursor(g.CurrentCell.X + 1, g.CurrentCell.Y + 1)

			if g.IsClear() {
				g.ShowMessage("Congratulations!", term.ColorYellow)
				break Loop
			}

			switch term.Getc() {
			case 'q':
				g.ShowMessage("bye bye!", term.ColorBlack)
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
					g.ShowMessage(err.Error(), term.ColorRed)
					break Loop
				}
			case 'f':
				if !g.CurrentCell.IsOpened {
					g.CurrentCell.ToggleDangerSign()
					if g.CurrentCell.DangerSign {
						g.DangerSignCount += 1
					} else {
						g.DangerSignCount -= 1
					}
				}
			}

		}
	})
}
