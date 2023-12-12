package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
