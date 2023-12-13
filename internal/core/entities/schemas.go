package entities

var ProjectListSchema = `
{	
	"type": "array",
	"items": {
		"type": "object",
		"properties": {
			"name": {
				"type":"string"
			},
			"url": {
				"type":"string"
			},
			"type": {
				"type":"string",
				"enum": ["github-repo", "github-gist"]
			},
			"scripts": {
				"type":"object",
				"properties": {
					"install_script": {
						"type": "string"
					}
				},
				"required": ["install_script"]
			}
		},
		"required": ["name","url","type"]
	}
}`
