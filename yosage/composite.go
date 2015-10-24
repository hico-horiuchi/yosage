package yosage

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
)

func composite(input *gif.GIF, lgtm image.Image) (*gif.GIF, error) {
	output := gif.GIF{
		Delay:     input.Delay,
		LoopCount: input.LoopCount,
		Disposal:  input.Disposal,
		Config: image.Config{
			Width:  input.Config.Width,
			Height: input.Config.Height,
		},
	}

	x := input.Config.Width/2 - lgtm.Bounds().Dx()/2
	y := input.Config.Height/2 - lgtm.Bounds().Dy()/2

	fmt.Print("compositting frame...")

	for i, frame := range input.Image {
		fmt.Printf("%d ", i+1)

		draw.Draw(frame, frame.Bounds(), lgtm, image.Point{-x, -y}, draw.Over)
		output.Image = append(output.Image, frame)
	}

	fmt.Println("done")
	return &output, nil
}
