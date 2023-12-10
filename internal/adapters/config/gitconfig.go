package config

import (
	"errors"
	"github.com/realfabecker/bogo/internal/core/ports"
	"io/ioutil"
	"net/http"
)

// GistRepoConfigDownloader gist config struct definition
type GistRepoConfigDownloader struct {
	url string
}

// NewGistRepoConfigDownloader config service constructor
func NewGistRepoConfigDownloader(url string) ports.RepoConfigDownloader {
	return &GistRepoConfigDownloader{url}
}

// Download obtain repository configration
func (g GistRepoConfigDownloader) Download() ([]byte, error) {
	r, err := http.Get(g.url)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.New("unable to download repo config")
	}
	return data, nil
}
