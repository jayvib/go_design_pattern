package shapes

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
)

type ImageSquare struct {
	Output // interface embedding to satisfy the StrategyV2 interface.
}

func (i *ImageSquare) Print() error {
	width := 800
	height := 600

	origin := image.Point{0, 0} // the starting point of the top left end of the image

	bgImage := image.NewRGBA( // Create an rectangle image
		image.Rectangle{
			Min: origin,
			Max: image.Point{X: width, Y: height},
		},
	)

	bgColor := image.Uniform{ // create a color object
		C: color.RGBA{R: 70, G: 70, B: 70, A: 0},
	}
	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src) // draw an image

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{
		C: color.RGBA{R: 255, G: 0, B: 0, A: 1},
	}

	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})

	squareImg := image.NewRGBA(square)
	draw.Draw(bgImage, squareImg.Bounds(), &squareColor, origin, draw.Src)
	quality := &jpeg.Options{Quality: 75}

	if err := jpeg.Encode(i, bgImage, quality); err != nil {
		return err
	}

	_, err := i.LogWrite([]byte("Image written in provided writer\n"))
	if err != nil {
		return err
	}
	return nil
}
