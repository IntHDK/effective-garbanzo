package database

import (
	"time"
)

func init() {
	init_errors()
}

type DatabaseModule interface {
	//Connection
	connect(DSN string, DBType string, Migrate bool) (result bool, err error)
	disconnect() (result bool, err error)
	IsReady() (ready bool)

	//Query
	SearchPostList(Keyword_Title string, Keyword_Author string,
		CreateAt_Start time.Time, CreateAt_End time.Time,
		Size int, Offset int, Sort []struct {
			SortBy      string
			IsAscending bool
		}) (totalcount int64, result []ModelPostListRecord, err error)
	GetPost(UUID string) (result ModelPost, err error)
	AddPost(Source ModelPost, Password string) (err error)
	UpdatePost(Source ModelPost, Password string) (err error)
	DeletePost(UUID string, Password string) (err error)
}
