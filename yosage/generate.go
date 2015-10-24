package yosage

import "errors"

func Generate(inputPath, outputPath string) error {
	input, err := readInput(inputPath)
	if err != nil {
		return err
	}

	lgtm, err := readLGTM()
	if err != nil {
		return err
	}

	if input.Config.Width < lgtm.Bounds().Dx() || input.Config.Height < lgtm.Bounds().Dy() {
		return errors.New("yosage: input gif is smaller than LGTM png")
	}

	output, err := composite(input, lgtm)
	if err != nil {
		return err
	}

	output.LoopCount = 0
	return write(outputPath, output)
}
