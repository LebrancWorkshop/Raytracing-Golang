package template

import (
	"strconv"
)

type Template struct {
	Format string
	Width  int
	Height int
	MaxColor int
}

var template *Template

func PPM(format string, width int, height int, maxcolor int) *Template {
	template = &Template{
		Format: format,
		Width:  width,
		Height: height,
		MaxColor: maxcolor,
	}
	return template
}

func (t *Template) Starter() string {
	widthFormat := strconv.Itoa(t.Width)
	heightFormat := strconv.Itoa(t.Height)
	maxcolorFormat := strconv.Itoa(t.MaxColor)
	header := t.Format + "\n" + widthFormat + " " + heightFormat + "\n" + maxcolorFormat + "\n"
	return header
}