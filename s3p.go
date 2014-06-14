package main

import (
	"fmt"
	"github.com/mtharrison/s3p/files"
	"github.com/mtharrison/s3p/settings"
	"log"
)

func main() {

	// Parse the settings given to the command
	settings, err := settings.ParseSettings()

	if err != nil {
		log.Fatal(err)
	}

	files, err := files.GetFiles(settings.SourcePath)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)

}
