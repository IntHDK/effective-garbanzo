package webserver

import (
	"log"
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

	logger *log.Logger
}

type WebAllServerConfiguration struct {
	ListenAt string
	Logger   *log.Logger
}

func NewWebAllServer(configuration WebAllServerConfiguration) (server *WebAllServer) {
	server = &WebAllServer{
		httpserver:       &http.Server{},
		router:           &mux.Router{},
		ws_upgrader:      &websocket.Upgrader{},
		configuration:    configuration,
		wssessionmanager: NewWebsocketSessionManager(),
		logger:           configuration.Logger,
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
