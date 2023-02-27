package main

import (
	"go-raytracing/draw"
	"go-raytracing/file"
	"go-raytracing/template"
	"go-raytracing/config"
)

func main() {
	template := template.PPM(config.FORMAT, config.COLUMNS, config.ROWS, config.MAX_COLOR)
	header := template.Starter()
	content := draw.Draw()
	image := header + "\n" + content

	file.WriteFile("image", image)
}