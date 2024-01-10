package webserver

import (
	_ "embed"
	"html/template"
	"sync"
)

var templates sync.Map

func init() {
	templates.Store("http_view_index", parsetemplate("http_view_index", source_http_view_index))
}
func parsetemplate(name string, context string) (res *template.Template) {
	res, _ = template.New(name).Parse(context)
	//에러처리 추가
	return
}

// Template sources
//
//go:embed template/http_view_index.html
var source_http_view_index string
