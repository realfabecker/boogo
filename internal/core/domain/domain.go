package domain

type ProjectType string

const (
	TypeGithubGist ProjectType = "github-gist"
	TypeGithubRepo ProjectType = "github-repo"
)

type ProjectVarType string

const (
	VarTypeStdin ProjectVarType = "stdin"
)

type ProjectScript struct {
	Install *string `yaml:"install,omitempty" json:"install,omitempty"`
}

type Project struct {
	Name        string                 `yaml:"name,omitempty" json:"name,omitempty"`
	Alias       string                 `yaml:"alias,omitempty" json:"alias,omitempty"`
	Description string                 `yaml:"description,omitempty" json:"description,omitempty"`
	Url         string                 `yaml:"url,omitempty" json:"url,omitempty"`
	Type        ProjectType            `yaml:"type,omitempty" json:"type,omitempty"`
	Scripts     *ProjectScript         `yaml:"scripts,omitempty" json:"scripts,omitempty"`
	Vars        map[string]*ProjectVar `yaml:"vars,omitempty" json:"vars,omitempty"`
}

func (p *Project) GetUse() string {
	if p.Alias != "" {
		return p.Alias
	}
	return p.Name
}

type ProjectVar struct {
	Type        ProjectVarType `yaml:"type,omitempty" json:"type,omitempty"`
	Description string         `yaml:"description,omitempty" json:"description,omitempty"`
	Value       string         `yaml:"value,omitempty" json:"value,omitempty"`
}

type RepoConfig struct {
	Projects []Project `yaml:"projects" json:"projects"`
}

type BogoConfig struct {
	RepoUrl  string `json:"repo_url,omitempty"`
	RepoAuth string `json:"repo_auth,omitempty"`
}
