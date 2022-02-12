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
		Rows:              8,
		Columns:           3,
		MarginWidth:       1,
		LineStrokeWidth:   0.2,
		BorderStrokeWidth: 0.4,
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.DrawImage(0, 1, gridder.ImageConfig1{Color: color.Black, File: "./bincho.png"})
	grid.DrawImage(1, 0, gridder.ImageConfig1{Color: color.Black, File: "./bincho.png"})
	grid.DrawImage(1, 2, gridder.ImageConfig1{Color: color.Black, File: "./bincho.png"})
	grid.DrawImage(2, 1, gridder.ImageConfig1{Color: color.Black, File: "./bincho.png"})
	grid.DrawImage(3, 0, gridder.ImageConfig1{Color: color.Black, File: "./bincho.png"})
	grid.DrawImage(3, 2, gridder.ImageConfig1{Color: color.Black, File: "./bincho.png"})
	grid.DrawImage(4, 1, gridder.ImageConfig1{Color: color.Black, File: "./bincho.png"})
	grid.DrawImage(5, 0, gridder.ImageConfig1{Color: color.Black, File: "./bincho.png"})
	grid.DrawImage(5, 2, gridder.ImageConfig1{Color: color.Black, File: "./bincho.png"})
	grid.DrawImage(6, 1, gridder.ImageConfig1{Color: color.Black, File: "./bincho.png"})
	grid.DrawImage(7, 1, gridder.ImageConfig1{Color: color.Black, File: "./bincho.png"})

	grid.DrawCircle(0, 1, gridder.CircleConfig{Color: color.Black})
	grid.DrawCircle(1, 0, gridder.CircleConfig{Color: color.Black})
	grid.DrawCircle(1, 2, gridder.CircleConfig{Color: color.Black})
	grid.DrawCircle(2, 1, gridder.CircleConfig{Color: color.Black})
	grid.DrawCircle(3, 0, gridder.CircleConfig{Color: color.Black})
	grid.DrawCircle(3, 2, gridder.CircleConfig{Color: color.Black})
	grid.DrawCircle(4, 1, gridder.CircleConfig{Color: color.Black})
	grid.DrawCircle(5, 0, gridder.CircleConfig{Color: color.Black})
	grid.DrawCircle(5, 2, gridder.CircleConfig{Color: color.Black})
	grid.DrawCircle(6, 1, gridder.CircleConfig{Color: color.Black})
	grid.DrawCircle(7, 1, gridder.CircleConfig{Color: color.Black})

	grid.SavePNG()
}
