package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//var allEntries = make(map[string]string)
var sectionEntries = make(map[string]string)
var currentSection = ""

func ProcessCommands(args []string) {
	if args[0] == "check" {
		file, err := os.Open(args[1])
		check(err)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			text := scanner.Text()
			processLine(text)
		}
		defer func() {
			err = file.Close()
			check(err)
		}()
	}
}

func processLine(line string) {
	if strings.HasPrefix(line, "#") || len(line) == 0 {
		return
	}
	if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
		currentSection = line
		sectionEntries = make(map[string]string)
		return
	}
	entry := strings.Split(line, "=")
	if _, exists := sectionEntries[entry[0]]; exists {
		fmt.Printf("Section: %s => Duplicate line found: %s\n", currentSection, entry[0])
	} else {
		sectionEntries[entry[0]] = entry[1]
	}
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
