package yosage

import (
	"errors"
	"fmt"

	"github.com/gographics/imagick/imagick"
)

func readInput(inputPath string) (*imagick.MagickWand, error) {
	fmt.Print("reading input gif...")

	input := imagick.NewMagickWand()

	if err := input.ReadImage(inputPath); err != nil {
		fmt.Println()
		return input, err
	}

	fmt.Println("done")

	if input.GetImageFormat() != "GIF" {
		return input, errors.New("yosage: input is not gif")
	}

	return input, nil
}

var LGTM interface{}

func readLGTM() (*imagick.MagickWand, error) {
	fmt.Print("reading LGTM png...")

	var err error
	lgtm := imagick.NewMagickWand()

	switch LGTM.(type) {
	case string:
		err = lgtm.ReadImage(LGTM.(string))
		fmt.Println("done")
	case []byte:
		err = lgtm.ReadImageBlob(LGTM.([]byte))
		fmt.Println("done")
	}

	if err != nil {
		return lgtm, err
	} else if lgtm.GetImageFormat() != "PNG" {
		return lgtm, errors.New("yosage: LGTM is not png")
	}

	return lgtm, nil
}
