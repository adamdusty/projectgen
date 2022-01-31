package generator

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/adamdusty/projectgen/pkg/template"
)

func GenerateFile(file template.ProjectFile, stream io.Writer) error {
	return errors.New("Unimpl")
}

func GenerateProject(root string, proj template.RenderedTemplate) error {
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
