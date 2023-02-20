package cmd

import (
	"image/png"
	"os"

	"github.com/go-micah/imaging"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stepsCmd)
	stepsCmd.Flags().IntVar(&height, "height", 512, "height in pixels")
	stepsCmd.Flags().IntVar(&width, "width", 512, "width in pixels")
	stepsCmd.Flags().StringVarP(&filename, "filename", "f", "output.png", "desired output filename")
}

var stepsCmd = &cobra.Command{
	Use:   "steps",
	Short: "Generates a step wedge test image",
	Long: `Generates a step wedge test image.
Uses width and height parameters.
Saves the output as a PNG file`,

	Run: func(cmd *cobra.Command, args []string) {

		img := imaging.Steps(width, height)
		f, _ := os.Create(filename)
		png.Encode(f, img)
	},
}
