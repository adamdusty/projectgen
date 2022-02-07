package pgen

// TODO: Test template equality functions

type ProjectFile struct {
	Path    string
	Content string
}

type ProjectTemplate struct {
	Files       []ProjectFile
	Directories []string
	Variables   []string
}

type RenderedTemplate struct {
	Files       []ProjectFile
	Directories []string
}

func (r *RenderedTemplate) Equals(other *RenderedTemplate) bool {
	if len(r.Files) != len(other.Files) {
		return false
	}

	if len(r.Directories) != len(other.Directories) {
		return false
	}

	for i := range r.Files {
		if r.Files[i].Path != other.Files[i].Path {
			return false
		}

		if r.Files[i].Content != other.Files[i].Content {
			return false
		}
	}

	for i := range r.Directories {
		if r.Directories[i] != other.Directories[i] {
			return false
		}
	}

	return true
}

func (t *ProjectTemplate) Equals(other *ProjectTemplate) bool {
	if len(t.Files) != len(other.Files) || len(t.Directories) != len(other.Directories) || len(t.Variables) != len(other.Variables) {
		return false
	}

	for i := range t.Files {
		if t.Files[i].Path != other.Files[i].Path {
			return false
		}

		if t.Files[i].Content != other.Files[i].Content {
			return false
		}
	}

	for i := range t.Directories {
		if t.Directories[i] != other.Directories[i] {
			return false
		}
	}

	for i := range t.Variables {
		if t.Variables[i] != other.Variables[i] {
			return false
		}
	}

	return true
}
