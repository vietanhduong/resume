package home

import (
	"errors"
	"fmt"
	"github.com/vietanhduong/resume/pkg/cerrors"
	"github.com/vietanhduong/resume/pkg/utils/file"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
)

type service struct {
	resume *Resume
}

func NewService(path string) (*service, *cerrors.CError) {
	// parse
	resume, err := parse(path)
	if err != nil {
		return nil, err
	}
	// validate
	if err = ValidateResume(resume); err != nil {
		return nil, err
	}
	return &service{resume: resume}, nil
}

func parse(path string) (*Resume, *cerrors.CError) {
	// validate input path is exist
	if !file.IsExist(path) {
		return nil, cerrors.NewCError(http.StatusNotFound, errors.New(fmt.Sprintf("%s does not exist!", path)))
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, cerrors.NewCError(http.StatusInternalServerError, err)
	}

	// parse yaml
	var resume *Resume
	err = yaml.Unmarshal(content, &resume)
	if err != nil {
		return nil, cerrors.NewCError(http.StatusInternalServerError, err)
	}

	return resume, nil
}
