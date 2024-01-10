package webserver

import (
	"html/template"
	"log"
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

func (s *WebAllServer) controller_any_echo(w http.ResponseWriter, r *http.Request) {
	c, err := s.ws_upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
