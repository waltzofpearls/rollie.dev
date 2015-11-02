package libs

import (
	"bytes"
	"text/template"
	"time"
)

var tpls = map[string]*template.Template{
	"404":   newTemplate("views/layout.html", "views/404.html"),
	"index": newTemplate("views/layout.html", "views/index.html"),
}

var funcs = template.FuncMap{
	"moment": moment,
}

func moment(fmt string) string {
	return time.Now().Format(fmt)
}

func newTemplate(files ...string) *template.Template {
	return template.Must(template.New("*").Funcs(funcs).ParseFiles(files...))
}

func ExecuteTemplate(name string, data interface{}) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := tpls[name].ExecuteTemplate(&buf, "base", data)
	return &buf, err
}
