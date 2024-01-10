package webserver

import "net/http"

func (s *WebAllServer) route() {
	//s.router.HandleFunc()
	s.router.HandleFunc("/", s.controller_get_index).Methods(http.MethodGet)
	s.router.HandleFunc("/echo", s.controller_any_echo).Schemes("http", "https")
}
