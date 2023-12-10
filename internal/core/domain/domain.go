package domain

// ProjectType enum definition
type ProjectType string

const (
	TypeProject ProjectType = "project"
)

// ProjectScript struct definition
type ProjectScript struct {
	InstallScript *string `json:"install_script,omitempty"`
}

// Project struct definition
type Project struct {
	Name    string         `json:"name,omitempty"`
	Url     string         `json:"url,omitempty"`
	Type    ProjectType    `json:"type,omitempty"`
	Scripts *ProjectScript `json:"scripts,omitempty"`
}

// Config repo config struct definition
type Config struct {
	RepoUrl string `json:"repo_url,omitempty"`
}
