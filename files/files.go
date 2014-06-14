package files

import (
	"path/filepath"
	"os"
	"regexp"
)

type File struct {
	Name, Path string
}

func GetFiles(sourcePath string) (files []File, err error) {

	filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {

		matched, _ := regexp.MatchString("^[^\\.].+", path)

		if matched {
			files = append(files, File{info.Name(), path})
		}

		return err

	})

	return 

}