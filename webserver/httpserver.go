package webserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type WebAllServer struct {
	httpserver  *http.Server
	router      *mux.Router
	ws_upgrader *websocket.Upgrader

	configuration WebAllServerConfiguration

	wssessionmanager *websocketSessionManager
}

type WebAllServerConfiguration struct {
	ListenAt string
}

func NewWebAllServer(configuration WebAllServerConfiguration) (server *WebAllServer) {
	server = &WebAllServer{
		httpserver:       &http.Server{},
		router:           &mux.Router{},
		ws_upgrader:      &websocket.Upgrader{},
		configuration:    configuration,
		wssessionmanager: NewWebsocketSessionManager(),
	}

	//Setting router
	server.route()

	return
}

func (s *WebAllServer) HttpServerStart() error {
	s.httpserver.Addr = s.configuration.ListenAt
	s.httpserver.Handler = s.router
	return s.httpserver.ListenAndServe()
}
