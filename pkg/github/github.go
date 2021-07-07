package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vietanhduong/resume/pkg/cerrors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Github(user, repo, branch string) *github {
	return &github{
		user:   user,
		repo:   repo,
		branch: branch,
	}
}

// ghRawURL generate GitHub Raw URL
func (g *github) ghRawURL() string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", g.user, g.repo, g.branch)
}

// GetRaw make a request to GitHub to receive github file
func (g *github) GetRaw(path string) ([]byte, *cerrors.CError) {
	url := fmt.Sprintf("%s/%s", g.ghRawURL(), path)

	response, err := http.Get(url)
	if err != nil {
		return nil, cerrors.NewCError(http.StatusInternalServerError, err)
	}
	// make sure defer response body after made a request
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Fatalf("close response body error, err: %v", err)
		}
	}()
	// validate response code
	if response.StatusCode != http.StatusOK {
		return nil, cerrors.NewCError(response.StatusCode, errors.New(response.Status))
	}
	// read response body
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, cerrors.NewCError(http.StatusInternalServerError, err)
	}

	return content, nil
}

// SaveRaw write content to disk
func (g *github) SaveRaw(content []byte, output string) *cerrors.CError {
	err := ioutil.WriteFile(output, content, os.FileMode(0644))
	if err != nil {
		return cerrors.NewCError(http.StatusInternalServerError, err)
	}
	return nil
}

// CanPush validate input token can push the repo
// input token should look like:
// token <gh_token>
// or
// Bearer <gh_token>
func (g *github) CanPush(token string) (bool, *cerrors.CError) {
	// make sure the token can push this repo we can use https://api.github.com/repos
	// and permission block in response body
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", g.user, g.repo)
	// init client
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false, cerrors.NewCError(http.StatusInternalServerError, err)
	}
	// add header for request, i'm using github api v3
	request.Header = http.Header{
		"Accept":        []string{"application/vnd.github.v3+json"},
		"Authorization": []string{token},
	}
	// receive response
	response, err := client.Do(request)
	if err != nil {
		return false, cerrors.NewCError(http.StatusInternalServerError, err)
	}
	// make sure defer response body after made a request
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Fatalf("close response body error, err: %v", err)
		}
	}()
	// validate response code
	if response.StatusCode != http.StatusOK {
		return false, cerrors.NewCError(response.StatusCode, errors.New(response.Status))
	}
	// read response body
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, cerrors.NewCError(http.StatusInternalServerError, err)
	}
	// cast to RepoResponse
	var ghRepo RepoResponse
	if err := json.Unmarshal(content, &ghRepo); err != nil {
		return false, cerrors.NewCError(http.StatusInternalServerError, err)
	}
	if ghRepo.Permissions == nil || !ghRepo.Permissions.Push {
		return false, nil
	}
	return true, nil
}
