package files

import (
	"os"
	"path/filepath"
	"regexp"
)

type File struct {
	Name, Path string
	FileInfo   os.FileInfo
}

func GetFiles(sourcePath string, includeHidden bool) (files []File, err error) {

	filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {

		if includeHidden {
			files = append(files, File{info.Name(), path, info})
		} else {

			// Regex that matches files within a hidden path
			regex := "(^\\..+|.+\\/\\..+)"

			matched, _ := regexp.MatchString(regex, path)

			if !matched && !info.IsDir() {
				files = append(files, File{info.Name(), path, info})
			}
		}

		return err

	})

	return

}
