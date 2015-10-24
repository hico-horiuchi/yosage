package yosage

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkGIF(path string) error {
	r := regexp.MustCompile(`\.(gif|GIF)$`)
	if !r.MatchString(path) {
		return errors.New(fmt.Sprintf("yosage: %s is not gif", path))
	}

	return nil
}

func checkPNG(path string) error {
	r := regexp.MustCompile(`\.(png|PNG)$`)
	if !r.MatchString(path) {
		return errors.New(fmt.Sprintf("yosage: %s is not png", path))
	}

	return nil
}
