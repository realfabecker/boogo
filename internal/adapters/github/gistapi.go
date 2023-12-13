package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// Gist struct definition
type Gist struct {
	Id    string `json:"id"`
	Files map[string]struct {
		Filename string `json:"filename"`
		RawUrl   string `json:"raw_url"`
	}
}

// Api http client representation
type Api struct {
	token string
}

// NewApi definition of a new http client
func NewApi(token string) Api {
	return Api{token: token}
}

// GetGist obtain gist record from its id
func (g Api) GetGist(id string) (*Gist, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.github.com/gists/"+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", g.token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("body: %w", err)
	}

	if resp.StatusCode != 200 {
		fmt.Println(string(data))
		return nil, fmt.Errorf("status")
	}

	var gist Gist
	if err := json.Unmarshal(data, &gist); err != nil {
		return nil, fmt.Errorf("json: %w", err)
	}

	return &gist, nil
}

// Download dowloader function definition for app
func (g Api) Download(gist *Gist, dest string) error {
	for _, v := range gist.Files {
		if v.RawUrl == "" {
			return fmt.Errorf("lists: unable to list")
		}

		client := http.Client{}
		resp, err := client.Get(v.RawUrl)
		if err != nil {
			return fmt.Errorf("get:%w", err)
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("lists: unable to list")
		}

		if err := os.WriteFile(filepath.Join(dest, v.Filename), data, 0o644); err != nil {
			return fmt.Errorf("write-file: %s", err)
		}
	}
	return nil
}

// GetFile returns the file content from gist
func (g Api) GetFile(gist *Gist, file string) ([]byte, error) {
	v, ok := gist.Files[file]
	if !ok || v.RawUrl == "" {
		return nil, fmt.Errorf("lists: unable to list")
	}

	client := http.Client{}
	resp, err := client.Get(v.RawUrl)
	if err != nil {
		return nil, fmt.Errorf("get:%w", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("lists: unable to list")
	}
	return data, nil
}
