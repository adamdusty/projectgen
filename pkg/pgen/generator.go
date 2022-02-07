package pgen

import (
	"io/fs"
	"os"
	"path/filepath"
)

func GenerateProject(root string, proj RenderedTemplate) error {
	// Loop through and create directories
	for _, d := range proj.Directories {
		dir := filepath.Join(root, d)

		err := os.MkdirAll(dir, fs.ModeDir)
		if err != nil {
			return err
		}
	}

	// Loop through and create files
	for _, f := range proj.Files {
		path := filepath.Join(root, f.Path)

		file, err := os.Create(path)
		if err != nil {
			return err
		}

		file.WriteString(f.Content)
		file.Close()
	}

	return nil
}
