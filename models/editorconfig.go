package models

type EditorConfig struct {
	IsRoot         bool
	DefaultSection EditorConfigSection
	Sections       map[EditorConfigSectionHeader]EditorConfigSection
}

func (e *EditorConfig) FindFirstEntry(key string) (EditorConfigEntry, bool) {
	for header := range e.Sections {
		section := e.Sections[header]
		for entry := range section.Entries {
			editorConfigEntry := section.Entries[entry]
			if editorConfigEntry.Key == key {
				return editorConfigEntry.Entry, true
			}
		}
	}
	return EditorConfigEntry{}, false
}

type EditorConfigSection struct {
	IsDefaultSection bool
	Header           EditorConfigSectionHeader
	FileExpressions  []string
	Entries          map[string]EditorConfigEntries
}

func (s *EditorConfigSection) FindEntry(key string) (EditorConfigEntry, bool) {
	entry, found := s.Entries[key]
	return entry.Entry, found
}

type EditorConfigSectionHeader string

type EditorConfigEntries struct {
	Key          string
	Entry        EditorConfigEntry
	otherEntries []EditorConfigEntry
}

func (entries *EditorConfigEntries) AddEntry(entry EditorConfigEntry) {
	entries.otherEntries = append(entries.otherEntries, entry)
	entries.Entry = entry
}

func (entries EditorConfigEntries) HasMultipleEntries() bool {
	return len(entries.otherEntries) > 0
}

type EditorConfigEntry struct {
	Key   string
	Value string
}
