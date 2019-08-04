package routes

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/waltzofpearls/rollie.dev/libs"
)

type Index struct {
	libs.Subrouter
}

func (sr *Index) AttachHandlers() {
	sr.Router.Handle("/", libs.HandlerFunc(sr.homeHandler)).Methods("GET")
	sr.Router.Handle("/{redirect:projects}", libs.HandlerFunc(sr.redirectHandler)).Methods("GET")
	sr.Router.Handle("/{redirect:resume}", libs.HandlerFunc(sr.redirectHandler)).Methods("GET")
}

func (sr *Index) homeHandler(w http.ResponseWriter, r *http.Request) error {
	return sr.HtmlResponseHandler(w, r, "index", map[string]interface{}{
		"Config": sr.Config,
		"Title":  "Rollie Ma - Polyglot Developer from Vancouver, BC",
		"Description": strings.Join([]string{
			"Hey, I'm Rollie:",
			"LEGO bricks and Linux enthusiast.",
			"Self-motivated and fascinated by robotics, computer vision and machine learning.",
			"Polyglot developer captivated by Go, Python and JavaScript.",
		}, " "),
		"Url":   "https://rollie.dev",
		"Image": "https://rollie.dev/images/logos/logo-120x120.png",
	})
}

func (sr *Index) redirectHandler(w http.ResponseWriter, r *http.Request) error {
	re := regexp.MustCompile("^/?")
	url := re.ReplaceAllLiteralString(r.URL.String(), "/#")
	return sr.RedirectHandler(w, r, url)
}
