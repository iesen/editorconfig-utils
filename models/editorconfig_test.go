package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditorConfig_FindFirstEntry_FindsIfExists(t *testing.T) {
	// Arrange
	expected := EditorConfigEntry{
		Key:   "indent_size",
		Value: "4",
	}
	editorConfig, _ := NewEditorConfig("../test/data/editorconfig/find_first_exists.txt")
	// Act
	entry, found := editorConfig.FindFirstEntry("indent_size")
	// Assert
	assert.True(t, found)
	assert.Equal(t, expected, entry)
}

func TestEditorConfig_FindFirstEntry_NotFoundIfNotExists(t *testing.T) {
	// Arrange
	editorConfig, _ := NewEditorConfig("../test/data/editorconfig/find_first_exists.txt")
	// Act
	entry, found := editorConfig.FindFirstEntry("some_strange_prop")
	// Assert
	assert.False(t, found)
	assert.Empty(t, entry)
}

func TestEditorConfigSection_FindEntry_FindsIfExists(t *testing.T) {
	// Arrange
	expected := EditorConfigEntry{
		Key:   "indent_style",
		Value: "tab",
	}
	editorConfig, _ := NewEditorConfig("../test/data/editorconfig/section_find_exists.txt")
	// Act
	entry, found := editorConfig.DefaultSection.FindEntry("indent_style")
	// Assert
	assert.True(t, found)
	assert.Equal(t, expected, entry)
}

func TestEditorConfigSection_FindEntry_DoesNotFindIfNotExists(t *testing.T) {
	// Arrange
	editorConfig, _ := NewEditorConfig("../test/data/editorconfig/section_find_exists.txt")
	// Act
	entry, found := editorConfig.DefaultSection.FindEntry("some_check")
	// Assert
	assert.False(t, found)
	assert.Empty(t, entry)
}

func TestEditorConfigEntries_AddEntry_ChangesValueAndPushesPreviousValueToAnotherField(t *testing.T) {
	// Arrange
	entries := EditorConfigEntries{
		Key: "mykey",
		Entry: EditorConfigEntry{
			Key:   "mykey",
			Value: "value1",
		},
		otherEntries: nil,
	}
	// Act
	entries.AddEntry(EditorConfigEntry{
		Key:   "mykey",
		Value: "value2",
	})
	// Assert
	assert.True(t, entries.HasMultipleEntries())
	assert.Equal(t, entries.Entry.Value, "value2")
}
