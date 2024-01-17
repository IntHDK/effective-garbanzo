package webserver

import (
	"net/http"
)

func (s *WebAllServer) controller_any_echo(w http.ResponseWriter, r *http.Request) {
	c, err := s.ws_upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			s.logger.Println("read:", err)
			break
		}
		s.logger.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			s.logger.Println("write:", err)
			break
		}
	}
}
