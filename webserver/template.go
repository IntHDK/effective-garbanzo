package webserver

import (
	_ "embed"
	"html/template"
	"sync"
)

var templates sync.Map

func init() {
	//declare
	templatenames := []string{
		"http_view_index",
	}

	for _, v := range templatenames {
		tmp, err := parsetemplate(v, source_http_view_index)
		if err == nil {
			templates.Store(v, tmp)
		}
	}
}
func parsetemplate(name string, context string) (res *template.Template, err error) {
	res, err = template.New(name).Parse(context)
	//에러처리 추가
	return
}

// Template sources
//
//go:embed template/http_view_index.html
var source_http_view_index string
