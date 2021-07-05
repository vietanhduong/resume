package github

import (
	"fmt"
	"github.com/vietanhduong/resume/pkg/cerrors"
	"io/ioutil"
	"log"
	"net/http"
)

type raw struct {
	user string
	repo string
}

func NewRaw(user, repo string) *raw {
	return &raw{
		user: user,
		repo: repo,
	}
}

func (r *raw) baseURL() string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s", r.user, r.repo)
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
