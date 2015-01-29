package main

import (
	"flag"
	"./minesweeper"
)

func main() {
	bomb := flag.Int("b", 100, "number of bombs")
	width := flag.Int("w", 40, "width")
	height := flag.Int("h", 20, "height")
	template := flag.String("t", "", "template: heart")
	flag.Parse()

	g := minesweeper.NewGame()
	if *template == "heart" {
		g.SetTemplateMap(minesweeper.HeartMap, *bomb)
	} else {
		g.SetRectangleMap(*width, *height, *bomb)
	}
	g.Start()
}
