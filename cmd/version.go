package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.0.1"

func newVersionCommand() *cobra.Command {
	version := &cobra.Command{
		Use:   "version",
		Short: "Prints version info",
		Long:  "Prints version info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("{\"version\": \"%s\"}", version)
		},
	}
	return version
}
