package main

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("CRONO_GITHUB_TOKEN")
	if len(token) < 1 {
		log.Fatal("Please set CRONO_GITHUB_TOKEN to a valid token in order to authenticate to github.")
	}
	client := NewClient(token)
	client.getRepositoryList()
}

// GitClient is a client with an authenticated http client a context
// and a github.Client.
type GitClient struct {
	ctx     context.Context
	client  *http.Client
	gClient *github.Client
}

var gClient *GitClient

// NewClient creates a new Archiver client. This will return an existing client
// if it already was created.
func NewClient(token string) *GitClient {
	log.Println("Creating github client")
	if gClient != nil {
		return gClient
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	gClient = &GitClient{
		ctx:     ctx,
		client:  tc,
		gClient: client,
	}
	return gClient
}

func (client GitClient) getRepositoryList() []*github.Repository {
	user, _, _ := client.gClient.Users.Get(client.ctx, "")
	log.Println("Client created...")
	// list all repositories for the authenticated user
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	var allRepos []*github.Repository
	log.Println("Gathering repositories...")
	for {
		repos, resp, err := client.gClient.Repositories.List(client.ctx, user.GetLogin(), opt)
		if err != nil {
			log.Fatal("error retrieving repositories: ", err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	log.Printf("Retrieved %d repositories.\n", len(allRepos))
	return allRepos
}
