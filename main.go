package main

import (
	"go-raytracing/draw"
	"go-raytracing/file"
	"go-raytracing/template"
)

const (
	FORMAT = "P3"
	COLUMNS = 3
	ROWS = 2
	MAX_COLOR = 255
)

func main() {
	template := template.PPM("P3", 3, 2, 255)
	header := template.Starter()
	content := draw.Draw()
	image := header + "\n" + content

	file.WriteFile("image", image)
}