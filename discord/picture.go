package discord

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type Picture interface {
	Name() string
	Body() io.Reader
}

func NewFSPicture(path string) (Picture, error) {
	if path == "" {
		return nil, errors.New("path should not be empty")
	}

	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file(%s): %w", path, err)
	}
	return &FilePicture{
		path:   path,
		reader: bytes.NewReader(fileContent),
	}, nil
}

type FilePicture struct {
	path   string
	reader io.Reader
}

func (p FilePicture) Name() string {
	return filepath.Base(p.path)
}

func (p FilePicture) Body() io.Reader {
	return p.reader
}

func NewURLPicture(name, url string) (Picture, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get image: %w", err)
	}

	defer response.Body.Close()
	urlContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read image from url: %w", err)
	}

	return &URLPicture{
		url:    url,
		name:   name,
		reader: bytes.NewReader(urlContent),
	}, nil
}

type URLPicture struct {
	url    string
	name   string
	reader io.Reader
}

func (u URLPicture) Name() string {
	return u.name
}

func (u URLPicture) Body() io.Reader {
	return u.reader
}
