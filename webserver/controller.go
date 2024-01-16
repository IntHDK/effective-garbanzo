package webserver

import (
	"html/template"
	"net/http"
)

func (s *WebAllServer) controller_get_index(w http.ResponseWriter, r *http.Request) {
	tm, ext := templates.Load("http_view_index")
	if ext {
		tmplate, ok := tm.(*template.Template)
		if ok {
			tmplate.Execute(w, "ws://"+r.Host+"/echo")
		}
	}
}
