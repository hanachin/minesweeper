package main

import (
	"./minesweeper"
)

func main() {
	g := minesweeper.NewGame()
	g.SetSquareMap(40, 20, 100)
	g.Start()
}
