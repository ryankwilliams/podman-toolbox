package imagecleaner

import (
	"fmt"
	"os"

	"github.com/ryankwilliams/podman-toolbox/pkg/podman"

	"github.com/spf13/cobra"

	"github.com/containers/podman/v4/pkg/bindings/images"
)

var Cmd = &cobra.Command{
	Use:   "image-cleaner",
	Short: "Clean unused images",
	Long:  "Clean unused images",
	Run:   run,
}

var flags struct {
	prompt                   bool
	removeDanglingImagesOnly bool
}

func init() {
	Cmd.Flags().BoolVarP(
		&flags.removeDanglingImagesOnly,
		"removeDanglingImagesOnly",
		"",
		false,
		"Remove only dangling images",
	)
	Cmd.Flags().BoolVarP(
		&flags.prompt,
		"prompt",
		"",
		false,
		"Interactive mode",
	)
}

func run(cmd *cobra.Command, argv []string) {
	connectionCtx, err := podman.CreateConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	localImages, err := images.List(connectionCtx, nil)
	if err != nil {
		fmt.Println("Failed to list images, err: %w", err)
		os.Exit(1)
	}

	removeImage := func(imageID string) error {
		var force = true
		_, err := images.Remove(connectionCtx, []string{imageID}, &images.RemoveOptions{Force: &force})
		if len(err) != 0 && err[0] != nil {
			return fmt.Errorf("failed to remove image: %s, error: %w", imageID, err[0])
		}
		return nil
	}

	for _, image := range localImages {
		if flags.prompt {
			var response string
			var imageName string

			if image.Dangling {
				imageName = "dangling-image"
			} else {
				imageName = image.Names[0]
			}

			// TODO: better handle user input
			fmt.Printf("Delete image: %s, (y|n)\n", imageName)
			fmt.Scan(&response)
			if response == "y" {
				err := removeImage(image.ID)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		} else if flags.removeDanglingImagesOnly && image.Dangling {
			err := removeImage(image.ID)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("Image: %s removed!\n", image.ID)
		} else if !flags.removeDanglingImagesOnly {
			err := removeImage(image.ID)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("Image: %s removed!\n", image.ID)
		}
	}
}
