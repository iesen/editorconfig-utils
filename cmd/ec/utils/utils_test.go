package utils

import "testing"

func TestProcessCommands(t *testing.T) {
	args := []string{"check", "../../../test/data/editorconfig_with_no_dups.txt"}
	ProcessCommands(args)
}
