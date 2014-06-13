package main

import (
	"os"
	"path/filepath"
	"fmt"
	"log"
	"regexp"
)

func main() {

	// Parse the settings given to the command
	settings, err := parseSettings()

	if err != nil {
		log.Fatal(err)
	}

	filepath.Walk(settings.SourcePath, func(path string, info os.FileInfo, err error) error {

		matched, _ := regexp.MatchString("^[^\\.].+", path)

		if matched {
			fmt.Println(path, ":", info.Name())
		}

		return err
	})

	// Get a list of the files and the folder heirarchy we need to upload

	// Start to upload the files one by one to the intended location

}
