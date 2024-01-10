package logic

import "time"

const (
	ERROR_DATASOURCE_NOTREADY       = "err_datasource_notready"
	ERROR_DATASOURCE_ENTITYNOTFOUND = "err_datasource_entitynotfound"
)

type Logic interface {
	SearchPostList(Keyword_Title string, Keyword_Author string,
		CreateAt_Start time.Time, CreateAt_End time.Time,
		Size int, Offset int, Sort []struct {
			SortBy      string
			IsAscending bool
		}) (totalcount int, result []Post, err error)
	GetPost(UUID string) (result Post, err error)
	AddPost(Source Post, Password string) (UUID string, err error)
	UpdatePost(Source Post, Password string) (err error)
	DeletePost(UUID string, Password string) (err error)
}

type Post struct {
	ID       string
	Author   string
	Title    string
	Context  string
	UpdateAt time.Time
}
