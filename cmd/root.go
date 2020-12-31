package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "ec",
		Aliases: []string{"editorconfigctl"},
		Short:   "Editorconfig utilities",
		Long:    "Editorconfig utilities that perform code tasks",
		Version: "0.0.1",
	}
	rootCmd.AddCommand(newVersionCommand())
	rootCmd.AddCommand(newCheckCommand())
	rootCmd.AddCommand(newMergeCommand())
	return rootCmd
}
