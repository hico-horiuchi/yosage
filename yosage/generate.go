package yosage

import (
	"errors"

	"github.com/gographics/imagick/imagick"
)

func Generate(inputPath, outputPath string) error {
	imagick.Initialize()
	defer imagick.Terminate()

	input, err := readInput(inputPath)
	defer input.Destroy()
	if err != nil {
		return err
	}

	lgtm, err := readLGTM()
	defer lgtm.Destroy()
	if err != nil {
		return err
	}

	if input.GetImageWidth() < lgtm.GetImageWidth() || input.GetImageHeight() < lgtm.GetImageHeight() {
		return errors.New("yosage: input gif is smaller than LGTM png")
	}

	output, err := composite(input, lgtm)
	defer lgtm.Destroy()
	if err != nil {
		return err
	}

	output.SetOption("loop", "0")
	write(outputPath, output)

	return nil
}
