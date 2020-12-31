package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditorConfig_Check_OkForValidFile(t *testing.T) {
	// Arrange
	expected := CheckResult{
		IsOk:     true,
		Messages: nil,
	}
	editorconfig, _ := NewEditorConfig("../test/data/editorconfig/check_ok.txt")
	// Act
	result, err := editorconfig.Check()
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestEditorConfig_Check_ErrorForNoCharsetOnDefaultSection(t *testing.T) {
	// Arrange
	expected := CheckResult{
		IsOk:     false,
		Messages: []string{"Default section should contain charset"},
	}
	editorconfig, _ := NewEditorConfig("../test/data/editorconfig/check_no_charset_on_default_section.txt")
	// Act
	result, err := editorconfig.Check()
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestEditorConfig_Check_CharsetValueIsNotUtf8(t *testing.T) {
	// Arrange
	expected := CheckResult{
		IsOk:     false,
		Messages: []string{"Default section charset value should be utf-8"},
	}
	editorconfig, _ := NewEditorConfig("../test/data/editorconfig/check_charset_is_not_utf8.txt")
	// Act
	result, err := editorconfig.Check()
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestEditorConfig_Check_NoDuplicateKeysWithinSectionNoError(t *testing.T) {
	// Arrange
	expected := CheckResult{
		IsOk:     true,
		Messages: nil,
	}
	editorconfig, _ := NewEditorConfig("../test/data/editorconfig/check_with_no_dups.txt")
	// Act
	result, err := editorconfig.Check()
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestEditorConfig_Check_WithDupsReturnsError(t *testing.T) {
	// Arrange
	expected := CheckResult{
		IsOk:     false,
		Messages: []string{"Default section contains duplicate config for key indent_size", "Default section contains duplicate config for key max_line_length"},
	}
	editorconfig, _ := NewEditorConfig("../test/data/editorconfig/check_with_dups.txt")
	// Act
	result, err := editorconfig.Check()
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
