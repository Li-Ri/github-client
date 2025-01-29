package github

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
)


type githubClient struct {
	GithubUser string
	BaseURL string
	Token string
	Client *http.Client
}


func NewGithubClient(baseURL, token string) *githubClient {
	return &githubClient{
		BaseURL: baseURL,
		Token: token,
		Client: &http.Client{},
	}
}


func (c *githubClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer " + c.Token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	res, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}


func (c *githubClient) GetCommitsByRepo(repo string) ([]github.CommitResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/repos/%s/%s/commits",c.GithubUser,repo), nil)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	resBody, err := c.Do(req)
	commits := []github.CommitResponse{}
	json.NewDecoder(resBody.Body).Decode(&commits)

	if(err != nil) {
		return nil, err
	}

	return commits, nil
}
