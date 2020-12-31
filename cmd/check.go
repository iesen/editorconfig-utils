package cmd

import (
	"fmt"

	"github.com/iesen/editorconfig-utils/models"
	"github.com/iesen/editorconfig-utils/utils"
	"github.com/spf13/cobra"
)

func newCheckCommand() *cobra.Command {
	checkCommand := &cobra.Command{
		Use:   "check",
		Short: "Checks given editorconfig file against common rules",
		Run: func(cmd *cobra.Command, args []string) {
			editorconfig, err := models.NewEditorConfig(args[0])
			utils.LogError(err)
			if err != nil {
				return
			}
			result, err := editorconfig.Check()
			utils.LogError(err)
			if err != nil {
				return
			}
			if result.IsOk {
				fmt.Printf("Check %s\nOK\n", args[0])
			} else {
				fmt.Printf("Check %s\nFAIL\n", args[0])
				fmt.Println("Messages:")
				for _, msg := range result.Messages {
					fmt.Printf("%s\n", msg)
				}
			}
		},
	}
	return checkCommand
}
