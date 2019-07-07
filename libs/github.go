package libs

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubRepo struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Fullname    *string `json:"fullname,omitempty"`
	Branch      *string `json:"branch,omitempty"`
	Url         *string `json:"url,omitempty"`
	Language    *string `json:"language,omitempty"`
	Forks       *int    `json:"forks,omitempty"`
	Stars       *int    `json:"stars,omitempty"`
	Watches     *int    `json:"watches,omitempty"`
	Badge       *string `json:"badge,omitempty"`
}

type GithubRepos []GithubRepo
type GithubContribs map[string]int
type docGetter func(*Config) (*goquery.Document, error)

func getGithubContribsDoc(c *Config) (*goquery.Document, error) {
	url := fmt.Sprintf(
		"https://github.com/users/%s/contributions",
		c.Github.Username,
	)
	return goquery.NewDocument(url)
}

type Github struct {
	Client    *github.Client
	DocGetter docGetter
	config    *Config
}

func NewGithub(config *Config) *Github {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Github.Token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return &Github{
		Client:    github.NewClient(tc),
		DocGetter: getGithubContribsDoc,
		config:    config,
	}
}

func (g *Github) GetRepos() (*GithubRepos, error) {
	opt := &github.RepositoryListOptions{
		Type:      "all",
		Sort:      "pushed",
		Direction: "desc",
	}
	repos, _, err := g.Client.Repositories.List(context.Background(), "", opt)
	if err != nil {
		return nil, err
	}

	var gr GithubRepos
	for _, repo := range repos {
		badge := fmt.Sprintf(
			"https://api.travis-ci.org/%s.svg?branch=%s",
			*repo.FullName,
			*repo.DefaultBranch,
		)
		gr = append(gr, GithubRepo{
			Name:        repo.Name,
			Description: repo.Description,
			Fullname:    repo.FullName,
			Branch:      repo.DefaultBranch,
			Url:         repo.HTMLURL,
			Language:    repo.Language,
			Forks:       repo.ForksCount,
			Stars:       repo.StargazersCount,
			Watches:     repo.SubscribersCount,
			Badge:       &badge,
		})
	}

	return &gr, nil
}

func (g *Github) GetContribs() (*GithubContribs, error) {
	doc, err := g.DocGetter(g.config)
	if err != nil {
		return nil, err
	}
	gc := make(GithubContribs)
	doc.Find("g > g > rect").Each(func(i int, s *goquery.Selection) {
		rawDate, _ := s.Attr("data-date")
		rawCount, _ := s.Attr("data-count")

		var (
			t         time.Time
			err       error
			timestamp string = "0"
			count     int    = 0
		)

		t, err = time.Parse("2006-01-02", rawDate)
		if err == nil {
			timestamp = strconv.FormatInt(t.Unix(), 10)
		}
		count, _ = strconv.Atoi(rawCount)

		gc[timestamp] = count
	})
	return &gc, nil
}
