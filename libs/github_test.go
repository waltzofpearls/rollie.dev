package libs_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/waltzofpearls/rolli3.net/libs"
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

var contribs string = `
<svg>
  <g>
    <g>
      <rect data-count="3" data-date="2014-11-21"/>
      <rect data-count="10" data-date="2014-11-22"/>
    </g>
    <g>
      <rect data-count="3" data-date="2014-11-23"/>
      <rect data-count="8" data-date="2014-11-24"/>
      <rect data-count="7" data-date="2014-11-29"/>
    </g>
    <g>
      <rect data-count="6" data-date="2014-11-30"/>
    </g>
  </g>
</svg>
`

var githubContribs string = `{
  "1416528000": 3,
  "1416614400": 10,
  "1416700800": 3,
  "1416787200": 8,
  "1417219200": 7,
  "1417305600": 6
}`

func TestGetRepos(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/user/repos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, repos)
	})

	cfg := libs.NewConfig()
	cfg.Github.Username = "u"
	cfg.Github.Token = "xxx"

	gh := libs.NewGithub(cfg)
	require.NotNil(t, gh)

	gh.Client = github.NewClient(nil)
	require.NotNil(t, gh.Client)

	url, _ := url.Parse(server.URL)
	gh.Client.BaseURL = url

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

func TestGetContribs(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, contribs)
		}),
	)
	defer server.Close()

	cfg := libs.NewConfig()
	gh := libs.NewGithub(cfg)
	require.NotNil(t, gh)

	gh.DocGetter = func(c *libs.Config) (*goquery.Document, error) {
		return goquery.NewDocument(server.URL)
	}

	gc, err := gh.GetContribs()
	require.Nil(t, err)
	require.NotNil(t, gc)

	b, err := json.Marshal(gc)
	require.Nil(t, err)
	require.NotEmpty(t, b)

	var dst bytes.Buffer
	json.Indent(&dst, b, "", "  ")

	assert.Equal(t, githubContribs, dst.String())
}
