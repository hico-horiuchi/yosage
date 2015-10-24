package yosage

import (
	"fmt"
	"image/gif"
	"os"
)

func write(outputPath string, output *gif.GIF) error {
	fmt.Print("writing...")

	if err := checkGIF(outputPath); err != nil {
		fmt.Println()
		return err
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = gif.EncodeAll(outputFile, output)
	if err != nil {
		return err
	}

	fmt.Println("done")
	return nil
}
