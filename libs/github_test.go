package libs

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-github/github"
)

func TestGetRepos(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	server.Close()

	url, _ := url.Parse(server.URL)

	cfg := NewConfig()
	cfg.Github.Token = "xxx"

	gh := NewGithub(cfg)
	gh.client = github.NewClient(nil)
	gh.client.BaseURL = url
}
