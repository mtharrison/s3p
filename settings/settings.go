package settings

import (
	"errors"
	"flag"
)

// CommandSettings holds the various available options for this command
type CommandSettings struct {
	DestinationPath string
	IncludeHidden   bool
	SourcePath      string
	BucketName      string
	AccessKeyID     string
	SecretAccessKey string
	Region          string
}

func (settings CommandSettings) isValid() bool {
	return settings.BucketName != "" && settings.AccessKeyID != "" && settings.SecretAccessKey != ""
}

func ParseSettings() (settings CommandSettings, err error) {

	flag.StringVar(&settings.DestinationPath, "dest", "/", "The destination path on S3, defaults to root")
	flag.StringVar(&settings.Region, "region", "us-east-1", "The region of the S3 bucket, defaults to us-east-1")
	flag.BoolVar(&settings.IncludeHidden, "includeHidden", false, "Include hidden files, off by default")
	flag.StringVar(&settings.SourcePath, "src", ".", "The source path on your local machine, defaults to '.'")
	flag.StringVar(&settings.BucketName, "bucket", "", "The destination path on S3, defaults to root")
	flag.StringVar(&settings.AccessKeyID, "keyid", "", "The destination path on S3, defaults to root")
	flag.StringVar(&settings.SecretAccessKey, "secretkey", "", "The destination path on S3, defaults to root")
	flag.Parse()

	if !settings.isValid() {
		err = errors.New("Error: You didn't give the correct settings. Type 's3p -help' for usage.")
	}

	return
}
