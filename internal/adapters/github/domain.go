package github

// Gist struct definition
type Gist struct {
	Id    string `json:"id"`
	Files map[string]struct {
		Filename string `json:"filename"`
		RawUrl   string `json:"raw_url"`
	}
}
