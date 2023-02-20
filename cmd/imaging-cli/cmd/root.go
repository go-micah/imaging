package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "imaging-cli",
	Short: "This program allows you to generate a number of test images.",
	Long:  `This program allows you to generate a number of test images.`,
}

var height, width int

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
