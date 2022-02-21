package pgen

import (
	"io/fs"
	"os"
	"path/filepath"
)

func GenerateProject(root string, project *ProjectTemplate, definitions map[string]interface{}) error {
	rendered, err := RenderTemplate(project, definitions)
	if err != nil {
		return err
	}

	for _, d := range rendered.Directories {
		dir := filepath.Join(root, d)

		err := os.MkdirAll(dir, fs.ModeDir)
		if err != nil {
			return err
		}
	}

	for _, f := range rendered.Files {
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
