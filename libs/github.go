package libs

import (
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubError struct {
}

type GithubRepo struct {
}

type Github struct {
	client *github.Client
	config *Config
}

func NewGithub(config *Config) *Github {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Github.Token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return &Github{
		client: github.NewClient(tc),
		config: config,
	}
}

func (g *Github) GetRepos() {
	opt := &github.RepositoryListOptions{
		Type:      "owner",
		Sort:      "pushed",
		Direction: "desc",
	}
	repos, _, err := g.client.Repositories.List(g.config.Github.Username, opt)
	if err != nil {
	}

	for _, repo := range repos {
		badge := fmt.Sprintf("https://api.travis-ci.org/%s.svg?branch=%s", repo.FullName, repo.DefaultBranch)
		log.Printf("%v", badge)
		// log.Printf("%v", repo)
	}
}
