package main

import (
	"fmt"
	"os"

	"github.com/iesen/editorconfig-utils/cmd"
)

func main() {
	err := cmd.NewRootCmd().Execute()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
