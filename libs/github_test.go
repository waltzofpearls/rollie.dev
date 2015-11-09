package libs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var repos string = `[
  {
    "id": 1,
    "name": "foo",
    "description": "bar",
    "branch": "master",
    "default_branch": "master",
    "full_name": "u",
    "html_url": "foobar",
    "language": "golang",
    "forks_count": 0,
    "stargazers_count": 0,
    "subscribers_count": 0
  }
]`

var githubRepos string = `[
  {
    "name": "foo",
    "description": "bar",
    "fullname": "u",
    "branch": "master",
    "url": "foobar",
    "language": "golang",
    "forks": 0,
    "stars": 0,
    "watches": 0,
    "badge": "https://api.travis-ci.org/u.svg?branch=master"
  }
]`

func TestGetRepos(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/users/u/repos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, repos)
	})

	cfg := NewConfig()
	cfg.Github.Username = "u"
	cfg.Github.Token = "xxx"

	gh := NewGithub(cfg)
	require.NotNil(t, gh)

	gh.client = github.NewClient(nil)
	require.NotNil(t, gh.client)

	url, _ := url.Parse(server.URL)
	gh.client.BaseURL = url

	gr, err := gh.GetRepos()
	require.Nil(t, err)
	require.NotNil(t, gr)

	b, err := json.Marshal(gr)
	require.Nil(t, err)
	require.NotEmpty(t, b)

	var dst bytes.Buffer
	json.Indent(&dst, b, "", "  ")

	assert.Equal(t, githubRepos, dst.String())
}
