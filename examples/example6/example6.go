package main

import (
	"image/color"
	"log"

	"github.com/shomali11/gridder"
)

func main() {
	imageConfig := gridder.ImageConfig{
		Width:  500,
		Height: 1000,
		Name:   "example6.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              7,
		Columns:           3,
		MarginWidth:       1,
		LineStrokeWidth:   0.4,
		BorderStrokeWidth: 0.4,
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.DrawImage(1, 1, gridder.ImageConfig1{Length: 1, Color: color.Black, File: "./bincho.png"})
	grid.DrawImage(2, 2, gridder.ImageConfig1{Length: 2, Color: color.Black, File: "./bincho.png", Rotate: 2})
	grid.DrawImage(1, 2, gridder.ImageConfig1{Length: 3, Color: color.Black, File: "./bincho.png"})
	grid.DrawImage(2, 1, gridder.ImageConfig1{Length: 40, Color: color.Black, File: "./bincho.png"})
	grid.SavePNG()
}
