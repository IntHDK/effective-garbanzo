package webserver

import (
	"effective-garbanzo/logic"
	"encoding/json"
	"math"
	"net/http"
	"time"
)

const (
	controller_get_post_list_request_Sortmode_None = iota
	controller_get_post_list_request_Sortmode_CreateAt_ASC
	controller_get_post_list_request_Sortmode_CreateAt_DESC
)
const (
	controller_get_post_list_PAGESIZE = 20
)

func (s *WebAllServer) controller_get_post_list(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Title               string
		Author              string
		IsUseFilterCreateAt bool
		Page                int
		Sortmode            int
	}
	errdecode := json.NewDecoder(r.Body).Decode(&request)
	if errdecode != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var st, en time.Time
	var offset int
	var sortcond []struct {
		SortBy      string
		IsAscending bool
	}

	switch request.Sortmode {
	case controller_get_post_list_request_Sortmode_None:
		fallthrough
	case controller_get_post_list_request_Sortmode_CreateAt_DESC:
		sortcond = append(sortcond, struct {
			SortBy      string
			IsAscending bool
		}{
			SortBy:      logic.PARAM_SearchPostList_SortBy_CreateAt,
			IsAscending: false,
		})
	case controller_get_post_list_request_Sortmode_CreateAt_ASC:
		sortcond = append(sortcond, struct {
			SortBy      string
			IsAscending bool
		}{
			SortBy:      logic.PARAM_SearchPostList_SortBy_CreateAt,
			IsAscending: true,
		})
	}

	offset = (request.Page - 1) * controller_get_post_list_PAGESIZE

	cnt, src, err := s.logicmodule.SearchPostList(request.Title, request.Author, st, en, controller_get_post_list_PAGESIZE, offset, sortcond)

	if err != nil {
		s.logger.Printf("controller_get_post_list error : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response struct {
		TotalPage  int
		TotalCount int64
		PostList   []struct {
			ID       string
			Title    string
			Author   string
			UpdateAt time.Time
		}
	}

	response.TotalCount = cnt
	response.TotalPage = int(math.Ceil(float64(cnt) / float64(controller_get_post_list_PAGESIZE)))
	response.PostList = []struct {
		ID       string
		Title    string
		Author   string
		UpdateAt time.Time
	}{}
	for _, v := range src {
		response.PostList = append(response.PostList, struct {
			ID       string
			Title    string
			Author   string
			UpdateAt time.Time
		}{
			ID:       v.ID,
			Title:    v.Title,
			Author:   v.Author,
			UpdateAt: v.UpdateAt,
		})
	}

	json.NewEncoder(w).Encode(response)

}

func (s *WebAllServer) controller_get_post(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ID string
	}
	errdecode := json.NewDecoder(r.Body).Decode(&request)
	if errdecode != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	src, err := s.logicmodule.GetPost(request.ID)
	if err != nil {
		s.logger.Printf("controller_get_post error : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response = struct {
		ID       string
		Title    string
		Context  string
		Author   string
		UpdateAt time.Time
	}{
		ID:       src.ID,
		Title:    src.Title,
		Context:  src.Context,
		Author:   src.Author,
		UpdateAt: src.UpdateAt,
	}

	json.NewEncoder(w).Encode(response)
}
