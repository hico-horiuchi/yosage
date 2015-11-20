package yosage

import (
	"errors"
	"image/gif"
)

func GenerateFromPath(inputPath, outputPath string) error {
	input, err := readInput(inputPath)
	if err != nil {
		return err
	}

	output, err := Generate(input)
	if err != nil {
		return err
	}

	return write(outputPath, output)
}

func Generate(input *gif.GIF) (*gif.GIF, error) {
	lgtm, err := readLGTM()
	if err != nil {
		return nil, err
	}

	if input.Config.Width < lgtm.Bounds().Dx() || input.Config.Height < lgtm.Bounds().Dy() {
		return nil, errors.New("yosage: input gif is smaller than LGTM png")
	}

	output, err := composite(input, lgtm)
	if err != nil {
		return nil, err
	}

	output.LoopCount = 0
	return output, nil
}
