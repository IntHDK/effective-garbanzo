package webserver

import "net/http"

func (s *WebAllServer) route() {
	//s.router.HandleFunc()
	s.router.HandleFunc("/", s.controller_get_index).Methods(http.MethodGet)
	s.router.HandleFunc("/ws/echo", s.controller_any_echo)

	s.router.HandleFunc("/post/list", s.controller_get_post_list).Methods(http.MethodGet)
}
