package domain

// ProjectType enum definition
type ProjectType string

const (
	TypeGithubGist ProjectType = "github-gist"
	TypeGithubRepo ProjectType = "github-repo"
)

// ProjectScript struct definition
type ProjectScript struct {
	InstallScript *string `json:"install_script,omitempty"`
}

// Project struct definition
type Project struct {
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Url         string         `json:"url,omitempty"`
	Type        ProjectType    `json:"type,omitempty"`
	Scripts     *ProjectScript `json:"scripts,omitempty"`
}

// Config repo config struct definition
type Config struct {
	RepoUrl  string `json:"repo_url,omitempty"`
	RepoAuth string `json:"repo_auth,omitempty"`
}
