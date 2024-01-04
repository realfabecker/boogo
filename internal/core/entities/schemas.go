package entities

var RepoConfigSchema = `
{
  "type": "object",
  "properties": {
    "projects": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "name": {
            "type": "string"
          },
          "url": {
            "type": "string"
          },
          "description": {
            "type": "string"
          },
          "type": {
            "type": "string",
            "enum": [
              "github-repo",
              "github-gist"
            ]
          },
		  "vars": {
 			 "type": "object",
			 "additionalProperties": {
				"type": "object",
				"properties": {
					"type": {
						"type": "string"
					},
					"description": {
						"type": "string"
					}
				}
			 }
 		  },
          "scripts": {
            "type": "object",
            "properties": {
              "install": {
                "type": "string"
              }
            },
            "required": [
              "install"
            ]
          }
        },
        "required": [
          "name",
          "url",
          "type",
          "description"
        ]
      }
    }
  }
}`
