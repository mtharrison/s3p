package settings

import (
	"errors"
	"flag"
)

// CommandSettings holds the various available options for this command
type CommandSettings struct {
	DestinationPath string
	SourcePath      string
	BucketName      string
	AccessKey       string
	AccessSecret    string
}

func (settings CommandSettings) isValid() bool {
	return settings.BucketName != "" && settings.AccessKey != "" && settings.AccessSecret != ""
}

func ParseSettings() (settings CommandSettings, err error) {

	flag.StringVar(&settings.DestinationPath, "dest", "/", "The destination path on S3, defaults to root")
	flag.StringVar(&settings.SourcePath, "src", ".", "The source path on your local machine, defaults to '.'")
	flag.StringVar(&settings.BucketName, "bucket", "", "The destination path on S3, defaults to root")
	flag.StringVar(&settings.AccessKey, "key", "", "The destination path on S3, defaults to root")
	flag.StringVar(&settings.AccessSecret, "secret", "", "The destination path on S3, defaults to root")
	flag.Parse()

	if !settings.isValid() {
		err = errors.New("Error: You didn't give the correct settings")
	}

	return
}
