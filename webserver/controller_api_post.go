package webserver

import (
	"encoding/json"
	"net/http"
	"time"
)

type controller_get_post_list_request struct {
	Title               string
	Author              string
	IsUseFilterCreateAt bool
	CreateAt_Start      time.Time
	CreateAt_End        time.Time
	Page                int
	Sortmode            int
}

const (
	controller_get_post_list_request_Sortmode_None = iota
	controller_get_post_list_request_Sortmode_CreateAt_ASC
	controller_get_post_list_request_Sortmode_CreateAt_DESC
)

func (s *WebAllServer) controller_get_post_list(w http.ResponseWriter, r *http.Request) {
	var request controller_get_post_list_request
	errdecode := json.NewDecoder(r.Body).Decode(&request)
	if errdecode != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
