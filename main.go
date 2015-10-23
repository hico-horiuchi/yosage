package main

import (

	// "github.com/gographics/imagick/imagick"
	"github.com/hico-horiuchi/yosage/yosage"
	"github.com/spf13/cobra"
)

var VERSION string

func main() {
	var (
		input   string
		output  string
		lgtm    string
		version bool
	)

	rootCmd := &cobra.Command{
		Use:   "yosage",
		Short: "Stabd-alone LGTM gif generator by Golang",
		Long:  "Stabd-alone LGTM gif generator by Golang\nhttps://github.com/hico-horiuchi/yosage",
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			if version {
				yosage.Version(VERSION)
			}

			if lgtm == "" {
				yosage.LGTM, err = Asset("lgtm.png")
				yosage.CheckError(err)
			} else {
				yosage.LGTM = lgtm
			}

			err = yosage.Generate(input, output)
			yosage.CheckError(err)
		},
	}

	rootCmd.Flags().StringVarP(&input, "input", "i", "input.gif", "input gif file path")
	rootCmd.Flags().StringVarP(&output, "output", "o", "output.gif", "output gif file path")
	rootCmd.Flags().StringVarP(&lgtm, "lgtm", "l", "", "LGTM text png file path")
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "print and check the version")

	rootCmd.Execute()
}
