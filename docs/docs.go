// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/render": {
            "get": {
                "description": "Render a OGP image with logo and text",
                "produces": [
                    "image/png"
                ],
                "summary": "Render a OGP image",
                "operationId": "render-ogp-image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Text to display(Needs space for automatic linebreak)",
                        "name": "text",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Web URL of the logo image to display(Show default image if error occured when loading)",
                        "name": "imgurl",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Background gradient start (top left) color(Color code in HEX without #)",
                        "name": "startcolor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Background gradient end (bottom right) color(Color code in HEX without #)",
                        "name": "endcolor",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "srvogimg",
	Description: "A Service that renders Open Graph Protocol image to share on social media",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
