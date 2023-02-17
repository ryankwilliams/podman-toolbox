package main

import (
	"os"

	"github.com/ryankwilliams/podman-toolbox/cmd/imagecleaner"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "podman-toolbox",
	Short: "Helpful commands used when working with podman",
	Long:  "Helpful commands used when working with podman",
	Run:   run,
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(imagecleaner.Cmd)
}

func run(cmd *cobra.Command, argv []string) {}
