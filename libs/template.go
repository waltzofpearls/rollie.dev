package libs

import (
	"bytes"
	"path/filepath"
	"text/template"
	"time"
)

type Template struct {
	path  string
	tpls  map[string]*template.Template
	funcs template.FuncMap
}

func NewTemplate(path string) *Template {
	tmpl := &Template{path: path}
	tmpl.funcs = template.FuncMap{
		"moment": tmpl.moment,
	}
	tmpl.tpls = map[string]*template.Template{
		"404":   tmpl.create("views/layout.html", "views/404.html"),
		"index": tmpl.create("views/layout.html", "views/index.html"),
	}
	return tmpl
}

func (t *Template) create(files ...string) *template.Template {
	for i, file := range files {
		files[i] = filepath.Join(t.path, file)
	}
	return template.Must(
		template.New("*").Funcs(t.funcs).ParseFiles(files...),
	)
}

func (t *Template) Execute(name string, data interface{}) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := t.tpls[name].ExecuteTemplate(&buf, "base", data)
	return &buf, err
}

func (t *Template) moment(format string) string {
	return time.Now().Format(format)
}
