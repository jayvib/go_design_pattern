package main

import "image"

type PrintStrategy interface {
	Print() error
}

type ConsoleSquare struct{}

func (ConsoleSquare) Print() error {
	println("square")
	return nil
}

type ImageSquare struct {
	destPath string
}

func (i *ImageSquare) Print() error {
	width := 800
	height := 600
	
	origin := image.Point{0, 0}

	bgImage := image.NewRGBA(
		image.Rectangle{
			Min: origin,
			Max: image.Point{X: width, Y: height},
		}
	)
	bgColor := image.Uniform{
		color.RGBA{
			R:70,
			G:70,
			B:70,
			A:0,
		}
	}
	quality := &jpeg.Options.{Quality:75}
	draw.Print(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)
	return nil
}
