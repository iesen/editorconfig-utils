package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newMergeCommand() *cobra.Command {
	var out string
	mergeCommand := &cobra.Command{
		Use:   "merge",
		Short: "Merges all files at given directory into final editorconfig file",
		Run: func(cmd *cobra.Command, args []string) {
			/*editorconfig, err := models.NewEditorConfig(args[0])
			utils.LogError(err)
			if err != nil {
				return
			}*/
			fmt.Println(args[0], out)
		},
	}
	mergeCommand.Flags().StringVar(&out, "out", ".editorconfig", "Result file name")
	return mergeCommand
}
