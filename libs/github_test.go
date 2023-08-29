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
	"github.com/waltzofpearls/rollie.dev/libs"
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
    "watches": 0
  }
]`

var contribs string = `
<table class="ContributionCalendar-grid js-calendar-graph-table">
  <tbody>
    <tr>
      <td class="ContributionCalendar-day" data-date="2014-11-21">
        <span>3 contributions on Friday, 21 Nov, 2014</span>
      </td>
      <td class="ContributionCalendar-day" data-date="2014-11-22">
        <span>10 contributions on Saturday, 22 Nov, 2014</span>
      </td>
    </tr>
    <tr>
      <td class="ContributionCalendar-day" data-date="2014-11-23">
        <span>3 contributions on Sunday, 23 Nov, 2014</span>
      </td>
      <td class="ContributionCalendar-day" data-date="2014-11-24">
        <span>8 contributions on Monday, 24 Nov, 2014</span>
      </td>
      <td class="ContributionCalendar-day" data-date="2014-11-29">
        <span>7 contributions on Tuesday, 29 Nov, 2014</span>
      </td>
    </tr>
    <tr>
      <td class="ContributionCalendar-day" data-date="2014-11-30">
        <span>6 contributions on Wednesday, 30 Nov, 2014</span>
      </td>
    </tr>
  </tbody>
</table>
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

	url, _ := url.Parse(server.URL + "/")
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
