package draw

import (
	"go-raytracing/config"
	"strconv"
)

func Draw() string {
	width := config.COLUMNS
	height := config.ROWS

	canvas := ""

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			canvas += strconv.Itoa(i) + " " + strconv.Itoa(j) + " " + strconv.Itoa(0) + "\n"
		}
	}

	return canvas
}