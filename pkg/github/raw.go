package github

import (
	"fmt"
	"github.com/vietanhduong/resume/pkg/cerrors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type raw struct {
	user   string
	repo   string
	branch string
}

func NewRaw(user, repo, branch string) *raw {
	return &raw{
		user:   user,
		repo:   repo,
		branch: branch,
	}
}

func (r *raw) baseURL() string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", r.user, r.repo, r.branch)
}

func (r *raw) GetRaw(path string) ([]byte, *cerrors.CError) {
	url := fmt.Sprintf("%s/%s", r.baseURL(), path)

	response, err := http.Get(url)
	if err != nil {
		return nil, cerrors.NewCError(http.StatusInternalServerError, err)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Fatalf("close response body error, err: %v", err)
		}
	}()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, cerrors.NewCError(http.StatusInternalServerError, err)
	}

	return content, nil
}

func (r *raw) SaveRaw(content []byte, output string) *cerrors.CError {
	err := ioutil.WriteFile(output, content, os.FileMode(0644))
	if err != nil {
		return cerrors.NewCError(http.StatusInternalServerError, err)
	}
	return nil
}
