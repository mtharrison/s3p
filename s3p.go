package main

import (
	"fmt"
	"github.com/mtharrison/s3p/auth"
	"github.com/mtharrison/s3p/files"
	"github.com/mtharrison/s3p/request"
	"github.com/mtharrison/s3p/settings"
	"log"
)

func main() {

	// Parse the settings given to the command
	settings, err := settings.ParseSettings()

	if err != nil {
		log.Fatal(err)
	}

	// Get the files
	files, err := files.GetFiles(settings.SourcePath,
		settings.IncludeHidden)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// Get the required headers
		headers := auth.GetHeaders(file, settings)

		// Do the actual request
		_, err := request.MakeRequest(file.Path, settings.DestinationPath, settings.BucketName, headers)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(file.Path, "=>", settings.BucketName+".amazonaws.com/"+settings.DestinationPath+file.Path, "copied ok")
	}

}
