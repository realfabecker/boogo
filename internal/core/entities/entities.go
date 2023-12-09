package entities

import "github.com/realfabecker/bogo/internal/core/domain"

func String(s string) *string {
	return &s
}

var Projects = map[string]domain.Project{
	"nodets": {
		Name: "nodets",
		Url:  "https://github.com/realfabecker/nodets",
		Type: domain.TypeProject,
		Scripts: &domain.ProjectScript{
			InstallScript: String("npm install"),
		},
	},
}
