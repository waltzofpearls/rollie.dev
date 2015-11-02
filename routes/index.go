package routes

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/waltzofpearls/tetris-go/libs"
)

type Index struct {
	libs.Subrouter
}

func (sr *Index) AttachHandlers() {
	sr.Router.Handle("/", libs.HandlerFunc(sr.homeHandler)).Methods("GET")
	sr.Router.Handle("/{redirect:(projects|resume)}", libs.HandlerFunc(sr.redirectHandler)).Methods("GET")
}

func (sr *Index) homeHandler(w http.ResponseWriter, r *http.Request) error {
	return sr.HtmlResponseHandler(w, r, "index", map[string]interface{}{
		"Config": sr.Config,
		"Title":  "Rollie Ma - Polyglot Developer from Vancouver, BC",
		"Description": strings.Join([]string{
			"Hi, I'm Rollie Ma. A Linux lover and LEGO bricks enthusiast.",
			"A polyglot developer obsessed with PHP, Python and JavaScript.",
			"A receptive learner captivated by mobile development, NoSQL and data mining.",
			"An amateur explorer interested in information aggregation and artificial intelligence fields.",
		}, " "),
		"Url":   "http://rolli3.net",
		"Image": "http://rolli3.net/images/logos/logo-120x120.png",
	})
}

func (sr *Index) redirectHandler(w http.ResponseWriter, r *http.Request) error {
	re := regexp.MustCompile("^/?")
	url := re.ReplaceAllLiteralString(r.URL.String(), "/#")
	return sr.RedirectHandler(w, r, url)
}
