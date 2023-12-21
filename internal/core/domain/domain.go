package domain

// ProjectType enum definition
type ProjectType string

const (
	TypeGithubGist ProjectType = "github-gist"
	TypeGithubRepo ProjectType = "github-repo"
)

// ProjectScript struct definition
type ProjectScript struct {
	Install *string `yaml:"install,omitempty" json:"install,omitempty"`
}

// Project struct definition
type Project struct {
	Name        string         `yaml:"name,omitempty" json:"name,omitempty"`
	Description string         `yaml:"description,omitempty" json:"description,omitempty"`
	Url         string         `yaml:"url,omitempty" json:"url,omitempty"`
	Type        ProjectType    `yaml:"type,omitempty" json:"type,omitempty"`
	Scripts     *ProjectScript `yaml:"scripts,omitempty" json:"scripts,omitempty"`
}

// RepoConfig projects list
type RepoConfig struct {
	Projects []Project `yaml:"projects" json:"projects"`
}

// BogoConfig repo config struct definition
type BogoConfig struct {
	RepoUrl  string `json:"repo_url,omitempty"`
	RepoAuth string `json:"repo_auth,omitempty"`
}
