package yosage

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"io"
	"os"
)

func readInput(inputPath string) (*gif.GIF, error) {
	fmt.Print("reading input gif...")

	if err := checkGIF(inputPath); err != nil {
		fmt.Println()
		return nil, err
	}

	inputFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Println()
		return nil, err
	}
	defer inputFile.Close()

	input, err := gif.DecodeAll(inputFile)
	if err != nil {
		fmt.Println()
		return nil, err
	}

	fmt.Println("done")
	return input, nil
}

var LGTM interface{}

func readLGTM() (image.Image, error) {
	var lgtmFile io.Reader
	fmt.Print("reading LGTM png...")

	switch LGTM.(type) {
	case string:
		lgtmPath := LGTM.(string)

		if err := checkPNG(lgtmPath); err != nil {
			fmt.Println()
			return nil, err
		}

		lgtmFile, err := os.Open(lgtmPath)
		if err != nil {
			fmt.Println()
			return nil, err
		}
		defer lgtmFile.Close()
	case []byte:
		lgtmFile = bytes.NewReader(LGTM.([]byte))
	}

	lgtm, err := png.Decode(lgtmFile)
	if err != nil {
		fmt.Println()
		return nil, err
	}

	fmt.Println("done")
	return lgtm, nil
}
