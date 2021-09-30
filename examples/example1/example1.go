package main

import (
	"image/color"
	"log"

	"github.com/shomali11/gridder"
)

func main() {
	imageConfig := gridder.ImageConfig{Name: "example1.png"}
	gridConfig := gridder.GridConfig{
		Rows:              7,
		Columns:           3,
		MarginWidth:       0,
		LineDashes:        1,
		LineStrokeWidth:   10,
		BorderDashes:      10,
		BorderStrokeWidth: 1,
		LineColor:         color.Gray{},
		BorderColor:       color.RGBA{B: 255, A: 255},
		BackgroundColor:   color.White,
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.SavePNG()
}
