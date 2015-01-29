package minesweeper

var HeartMap []string = []string{
	"  *********              *****   ",
	" *************        *********  ",
	" ***************    ************ ",
	"  ***************** ***********  ",
	"    **************************   ",
	"     ************************    ",
	"      *********************      ",
	"       ******************        ",
	"        **************           ",
	"        ***********              ",
	"       **********                ",
	"     ********                    ",
	"   *****                         ",
}

func NewTemplateMap(template []string) *Map {
	m := new(Map)

	cellCount := 0

	for _, line := range template {
		for _, c := range line {
			if c == '*' {
				cellCount += 1
			}
		}
	}

	m.Cells = make([]*Cell, cellCount)

	cells := make([][]Cell, len(template))
	for y := 0; y < len(template); y++ {
		cells[y] = make([]Cell, len(template[y]))
	}

	cellCount = 0

	rows := len(template)
	for y := 0; y < rows; y++ {
		cols := len(template[y])
		for x := 0; x < cols; x++ {
			c := &cells[y][x]
			c.X = x
			c.Y = y
			c.Neighbors = map[int]*Cell{}
			// up
			if (y - 1 >= 0 && template[y - 1][x] == '*') {
				c.Connect(DirUp, &cells[y - 1][x])
			}
			// up right
			if (y - 1 >= 0 && x + 1 < cols && template[y - 1][x + 1] == '*') {
				c.Connect(DirUpperRight, &cells[y - 1][x + 1])
			}
			// right
			if (x + 1 < cols && template[y][x + 1] == '*') {
				c.Connect(DirRight, &cells[y][x + 1])
			}
			// down right
			if (x + 1 < cols && y + 1 < rows && template[y + 1][x + 1] == '*') {
				c.Connect(DirLowerRight, &cells[y + 1][x + 1])
			}
			// down
			if (y + 1 < rows && template[y + 1][x] == '*') {
				c.Connect(DirBottm, &cells[y + 1][x])
			}
			// down left
			if (y + 1 < rows && x - 1 >= 0 && template[y + 1][x - 1] == '*') {
				c.Connect(DirLowerLeft, &cells[y + 1][x - 1])
			}
			// left
			if (x - 1 >= 0 && template[y][x - 1] == '*') {
				c.Connect(DirLeft, &cells[y][x - 1])
			}
			// left up
			if (x - 1 >= 0 && y - 1 >= 0 && template[y - 1][x - 1] == '*') {
				c.Connect(DirUpperLeft, &cells[y - 1][x - 1])
			}
			if template[y][x] == '*' {
				m.Cells[cellCount] = c
				cellCount += 1
			}
		}
	}

	return m
}
