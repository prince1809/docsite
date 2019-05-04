package docsite

import (
	"fmt"
	"github.com/pkg/errors"
	"html/template"
	"net/http"
	"path/filepath"
)

func parseTemplates(templateFS http.FileSystem, funcs template.FuncMap) (*template.Template, error) {
	tmpl := template.New("root")
	tmpl.Funcs(funcs)

	// Read all template files (recursively).
	isHTML := func(path string) bool { return filepath.Ext(path) == ".html" }
	err := WalkFileSystem(templateFS, isHTML, func(path string) error {
		data, err := ReadFile(templateFS, path)
		if err != nil {
			return errors.WithMessage(err, fmt.Sprintf("read template %s", path))
		}
		if _, err := tmpl.Parse(string(data)); err != nil {
			return errors.WithMessage(err, fmt.Sprintf("parse template %s", path))
		}
		return nil
	})
	return tmpl, errors.WithMessage(err, "walking templates")
}
