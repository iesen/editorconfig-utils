package models

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/iesen/editorconfig-utils/utils"
)

const defaultSectionKey = ""

var entryRegex = regexp.MustCompile(".+\\s*=\\s*.*")
var sectionRegex = regexp.MustCompile("\\[.+]")
var defaultSectionRegex = regexp.MustCompile("\\[\\s*\\*\\s*]")

func NewEditorConfig(filePath string) (EditorConfig, error) {
	file, err := os.Open(filePath)
	defer func() {
		err = file.Close()
		utils.LogError(err)
	}()
	if err != nil {
		utils.LogError(err)
		return EditorConfig{}, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defaultSection := EditorConfigSection{
		IsDefaultSection: true,
		Entries:          make(map[string]EditorConfigEntries),
	}
	editorConfig := EditorConfig{
		DefaultSection: defaultSection,
		Sections:       make(map[EditorConfigSectionHeader]EditorConfigSection),
	}
	editorConfig.Sections[defaultSectionKey] = defaultSection
	currentSection := editorConfig.DefaultSection
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if isComment(line) {
			continue
		}
		if sectionRegex.MatchString(line) {
			if isDefaultSectionHeader(line) {
				currentSection = editorConfig.DefaultSection
			} else {
				currentSection, err = editorConfig.newSection(line)
			}
			utils.LogError(err)
			if err != nil {
				return EditorConfig{}, err
			}
		} else if entryRegex.MatchString(line) {
			_, err := currentSection.newEntry(line)
			utils.LogError(err)
			if err != nil {
				return EditorConfig{}, err
			}
		} else {
			err := utils.CreateAndLogError(fmt.Sprintf("Invalid line in editorconfig: %s", line))
			return EditorConfig{}, err
		}
	}
	rootEntry, rootEntryFound := editorConfig.FindFirstEntry("root")
	editorConfig.IsRoot = rootEntryFound && rootEntry.Value == "true"
	return editorConfig, nil
}

func (e *EditorConfig) newSection(line string) (EditorConfigSection, error) {
	noSpaceLine := strings.ReplaceAll(line, " ", "")
	extensionsValue := strings.ReplaceAll(noSpaceLine, "[", "")
	extensionsValue = strings.ReplaceAll(extensionsValue, "]", "")
	section := EditorConfigSection{
		Header:          EditorConfigSectionHeader(noSpaceLine),
		FileExpressions: strings.Split(extensionsValue, ","),
		Entries:         make(map[string]EditorConfigEntries),
	}
	e.Sections[section.Header] = section
	return section, nil
}

func (s *EditorConfigSection) newEntry(line string) (EditorConfigEntry, error) {
	noSpacedLine := strings.ReplaceAll(line, " ", "")
	keyValue := strings.Split(noSpacedLine, "=")
	entry := EditorConfigEntry{
		Key:   keyValue[0],
		Value: keyValue[1],
	}
	existing, exists := s.Entries[entry.Key]
	if exists {
		existing.AddEntry(entry)
		s.Entries[entry.Key] = existing
	} else {
		s.Entries[entry.Key] = EditorConfigEntries{
			Key:   keyValue[0],
			Entry: entry,
		}
	}
	return entry, nil
}

func isDefaultSectionHeader(line string) bool {
	return defaultSectionRegex.MatchString(line)
}

func isComment(line string) bool {
	return strings.HasPrefix(line, "#")
}
