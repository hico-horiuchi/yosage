package yosage

import (
	"fmt"

	"github.com/gographics/imagick/imagick"
)

func write(outputPath string, output *imagick.MagickWand) error {
	fmt.Print("writing...")

	if err := output.WriteImages(outputPath, true); err != nil {
		fmt.Println()
		return err
	}

	fmt.Println("done")
	return nil
}
