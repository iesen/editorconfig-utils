package models

import (
	"fmt"
)

type CheckResult struct {
	IsOk     bool
	Messages []string
}

func (e EditorConfig) Check() (CheckResult, error) {
	charset, found := e.DefaultSection.FindEntry("charset")
	if !found {
		return CheckResult{
			IsOk:     false,
			Messages: []string{"Default section should contain charset"},
		}, nil
	}
	if charset.Value != "utf-8" {
		return CheckResult{
			IsOk:     false,
			Messages: []string{"Default section charset value should be utf-8"},
		}, nil
	}
	result, ok := checkDuplicateWithinSection(e)
	if !ok {
		return result, nil
	}
	return CheckResult{
		IsOk:     true,
		Messages: nil,
	}, nil
}

func checkDuplicateWithinSection(e EditorConfig) (CheckResult, bool) {
	messages := make([]string, 0)
	for header, section := range e.Sections {
		for key, entries := range section.Entries {
			if entries.HasMultipleEntries() {
				sectionName := header
				if section.IsDefaultSection {
					sectionName = "Default"
				}
				messages = append(messages, fmt.Sprintf("%s section contains duplicate config for key %s", sectionName, key))
			}
		}
	}
	if len(messages) > 0 {
		return CheckResult{
			IsOk:     false,
			Messages: messages,
		}, false
	}
	return CheckResult{}, true
}
