package commands

import (
	"image/png"
	"os"

	"github.com/go-micah/imaging"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(distCmd)
	distCmd.Flags().IntVar(&height, "height", 512, "height in pixels")
	distCmd.Flags().IntVar(&width, "width", 512, "width in pixels")
	distCmd.Flags().StringVarP(&filename, "filename", "f", "output.png", "desired output filename")
}

var distCmd = &cobra.Command{
	Use:   "dist",
	Short: "Generates a distance image",
	Long: `Generates a disatance image.
Uses with width and height parameters.
Saves the output as a PNG file`,

	Run: func(cmd *cobra.Command, args []string) {

		img := imaging.Dist(width, height)
		f, _ := os.Create(filename)
		png.Encode(f, img)
	},
}
