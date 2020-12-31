package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEditorConfig_FromRootOnlyConfig(t *testing.T) {
	// Arrange
	defaultSection := EditorConfigSection{
		IsDefaultSection: true,
		Header:           "",
		FileExpressions:  nil,
		Entries: map[string]EditorConfigEntries{
			"root": {
				Key: "root",
				Entry: EditorConfigEntry{
					Key:   "root",
					Value: "true",
				},
			},
			"charset": {
				Key: "charset",
				Entry: EditorConfigEntry{
					Key:   "charset",
					Value: "utf-8",
				},
			},
			"indent_size": {
				Key: "indent_size",
				Entry: EditorConfigEntry{
					Key:   "indent_size",
					Value: "4",
				},
			},
		},
	}
	expected := EditorConfig{
		IsRoot:         true,
		DefaultSection: defaultSection,
		Sections:       map[EditorConfigSectionHeader]EditorConfigSection{defaultSectionKey: defaultSection},
	}
	// Act
	editorConfig, err := NewEditorConfig("../test/data/factory/new_root_only_config.txt")
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, editorConfig)
}

func TestNewEditorConfig_FromEmptyFile(t *testing.T) {
	// Arrange
	defaultSection := EditorConfigSection{
		IsDefaultSection: true,
		Header:           "",
		FileExpressions:  nil,
		Entries:          make(map[string]EditorConfigEntries),
	}
	expected := EditorConfig{
		IsRoot:         false,
		DefaultSection: defaultSection,
		Sections:       map[EditorConfigSectionHeader]EditorConfigSection{defaultSectionKey: defaultSection},
	}
	// Act
	editorconfig, err := NewEditorConfig("../test/data/factory/new_empty_config.txt")
	//Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, editorconfig)
}

func TestNewEditorConfig_FromSectionedFile(t *testing.T) {
	// Arrange
	defaultSection := EditorConfigSection{
		IsDefaultSection: true,
		Header:           "",
		FileExpressions:  nil,
		Entries: map[string]EditorConfigEntries{
			"root": {
				Key: "root",
				Entry: EditorConfigEntry{
					Key:   "root",
					Value: "true",
				},
			},
			"indent_style": {
				Key: "indent_style",
				Entry: EditorConfigEntry{
					Key:   "indent_style",
					Value: "space",
				},
			},
			"indent_size": {
				Key: "indent_size",
				Entry: EditorConfigEntry{
					Key:   "indent_size",
					Value: "4",
				},
			},
		},
	}
	expected := EditorConfig{
		IsRoot:         true,
		DefaultSection: defaultSection,
		Sections: map[EditorConfigSectionHeader]EditorConfigSection{defaultSectionKey: defaultSection, "[*.go]": {
			IsDefaultSection: false,
			Header:           "[*.go]",
			FileExpressions:  []string{"*.go"},
			Entries: map[string]EditorConfigEntries{
				"indent_style": {
					Key: "indent_style",
					Entry: EditorConfigEntry{
						Key:   "indent_style",
						Value: "tab",
					},
				},
				"insert_final_newline": {
					Key: "insert_final_newline",
					Entry: EditorConfigEntry{
						Key:   "insert_final_newline",
						Value: "true",
					},
				},
			},
		}},
	}
	// Act
	editorconfig, err := NewEditorConfig("../test/data/factory/new_sectioned_config.txt")
	//Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, editorconfig)
}

func TestNewEditorConfig_NonRootEditorConfig(t *testing.T) {
	// Arrange
	defaultSection := EditorConfigSection{
		IsDefaultSection: true,
		Header:           "",
		FileExpressions:  nil,
		Entries: map[string]EditorConfigEntries{
			"charset": {
				Key: "charset",
				Entry: EditorConfigEntry{
					Key:   "charset",
					Value: "utf-8",
				},
			},
		},
	}
	expected := EditorConfig{
		IsRoot:         false,
		DefaultSection: defaultSection,
		Sections:       map[EditorConfigSectionHeader]EditorConfigSection{defaultSectionKey: defaultSection},
	}
	// Act
	editorconfig, err := NewEditorConfig("../test/data/factory/new_non_root_config.txt")
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, editorconfig)
}

func TestNewEditorConfig_InvalidConfigFileReturnsError(t *testing.T) {
	// Arrange
	// Act
	editorconfig, err := NewEditorConfig("../test/data/factory/new_invalid_config.txt")
	// Assert
	assert.Empty(t, editorconfig)
	assert.Error(t, err)
}
