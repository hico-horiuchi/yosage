package yosage

import (
	"fmt"

	"github.com/gographics/imagick/imagick"
)

func composite(input, lgtm *imagick.MagickWand) (*imagick.MagickWand, error) {
	output := imagick.NewMagickWand()

	x := int(input.GetImageWidth()/2 - lgtm.GetImageWidth()/2)
	y := int(input.GetImageHeight()/2 - lgtm.GetImageHeight()/2)

	fmt.Print("compositting frame...")

	for i := 0; i < int(input.GetNumberImages()); i++ {
		fmt.Printf("%d ", i+1)

		input.SetIteratorIndex(i)
		frame := input.GetImage()

		if err := frame.CompositeImage(lgtm, imagick.COMPOSITE_OP_OVER, x, y); err != nil {
			frame.Destroy()
			fmt.Println()
			return output, err
		}

		if err := output.AddImage(frame); err != nil {
			frame.Destroy()
			fmt.Println()
			return output, err
		}

		frame.Destroy()
	}

	fmt.Println("done")
	input.ResetIterator()

	return output, nil
}
