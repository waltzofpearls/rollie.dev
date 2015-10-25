package libs

import (
	"bytes"
	"net/http"
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

func ExecuteTemplate(w http.ResponseWriter, name string, status int, data interface{}) error {
	var buf bytes.Buffer

	err := tpls[name].ExecuteTemplate(&buf, "base", data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	w.Write(buf.Bytes())

	return nil
}
