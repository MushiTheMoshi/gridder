package main

import (
	"image/color"
	"log"

	"github.com/shomali11/gridder"
)

func main() {
	imageConfig := gridder.ImageConfig{
		Width:  500,
		Height: 500,
		Name:   "example6.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              4,
		Columns:           4,
		MarginWidth:       30,
		LineStrokeWidth:   2,
		BorderStrokeWidth: 4,
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.DrawImage(0, 0, gridder.ImageConfig1{Length: 1, Color: color.Black, File: "./bincho.png"})
	grid.DrawImage(0, 0, gridder.ImageConfig1{Length: 2, Color: color.Black, File: "./bincho.png", Rotate: 90})
	grid.DrawImage(0, 3, gridder.ImageConfig1{Length: 3, Color: color.Black, File: "./bincho.png", StrokeWidth: 25})
	grid.DrawImage(3, 3, gridder.ImageConfig1{Length: 4, Color: color.Black, File: "./bincho.png"})
	grid.SavePNG()
}
