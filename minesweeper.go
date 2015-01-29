package main

import (
	"flag"
	"./minesweeper"
)

func main() {
	bomb := flag.Int("b", 100, "number of bombs")
	width := flag.Int("w", 40, "width")
	height := flag.Int("h", 20, "height")
	flag.Parse()

	g := minesweeper.NewGame()
	g.SetSquareMap(*width, *height, *bomb)
	g.Start()
}
