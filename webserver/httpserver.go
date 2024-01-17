package webserver

import (
	"effective-garbanzo/logic"
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

	logicmodule      *logic.LogicModule
	wssessionmanager *websocketSessionManager

	logger *log.Logger
}

type WebAllServerConfiguration struct {
	ListenAt    string
	LogicModule *logic.LogicModule
	Logger      *log.Logger
}

func NewWebAllServer(configuration WebAllServerConfiguration) (server *WebAllServer) {
	server = &WebAllServer{
		httpserver:       &http.Server{},
		router:           &mux.Router{},
		ws_upgrader:      &websocket.Upgrader{},
		configuration:    configuration,
		wssessionmanager: NewWebsocketSessionManager(),
		logger:           configuration.Logger,
		logicmodule:      configuration.LogicModule,
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
