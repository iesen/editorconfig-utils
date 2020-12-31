package tasks

import (
	"github.com/iesen/editorconfig-utils/models"
)

/*
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/iesen/editorconfig-utils/models"
)
*/
type Task interface {
	Run(arguments []string) error
}

type TaskParameter struct {
}

var sections = make(map[string]models.EditorConfigSection)
var currentSectionEntries = make(map[string]string)
var currentSection = "root"

/*
func ProcessCommands(args []string) {
	command := args[1]
	if command == "check" {
		checkFile(args[1])
	} else if command == "merge" {
		mergeFiles(args[1])
		checkFile(args[1] + "/.editorconfig")
	} else if command == "compare" {
		checkFile(args[1])
		Compare(args)
	}
}

func checkFile(editorConfigFile string) {
	sections["root"] = models.EditorConfigSection{
		Name:    "",
		Entries: make(map[string]models.EditorConfigEntry),
	}
	file, err := os.Open(editorConfigFile)
	check(err)
	defer func() {
		err = file.Close()
		check(err)
	}()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		processLine(text)
	}
	checkForCharset()
}

func mergeFiles(folder string) {
	info, err := os.Stat(folder)
	check(err)
	var files []string
	result := ""
	if !info.IsDir() {
		log.Panicf("%s is not a directory\n", folder)
	}
	err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if info.Name() == ".editorconfig" {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	check(err)
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		check(err)
		result = result + string(data) + "\n"
	}
	err = ioutil.WriteFile(folder+"/.editorconfig", []byte(result), 0644)
	check(err)
	fmt.Printf("Merged editorconfig \"%s\" file from directory \"%s\"\n", folder+"/.editorconfig", folder)
}

func checkForCharset() {
	charset := sections["[*]"].Entries["charset"].Value
	if charset != "utf-8" {
		fmt.Printf("Non utf-8 charset found. Please fix this.\n")
	}

}

func processLine(line string) {
	if strings.HasPrefix(line, "#") || len(line) == 0 {
		return
	}
	if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
		currentSection = line
		currentSectionEntries = make(map[string]string)
		sections[line] = models.EditorConfigSection{
			Name:    line,
			Entries: make(map[string]models.EditorConfigEntry),
		}
		return
	}
	entry := strings.Split(line, "=")
	key := strings.TrimSpace(entry[0])
	value := strings.TrimSpace(entry[1])
	if _, exists := currentSectionEntries[key]; exists {
		fmt.Printf("Section: %s => Duplicate line found: %s\n", currentSection, key)
	} else {
		currentSectionEntries[key] = value
		sections[currentSection].Entries[key] = models.EditorConfigEntry{
			Key:   key,
			Value: value,
		}
	}
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
*/
