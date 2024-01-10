package database

import (
	"time"
)

func init() {
	init_errors()
}

type DatabaseModule interface {
	//Connection
	Connect(DSN string, DBType string, Migrate bool) (result bool, err error)
	Disconnect() (result bool, err error)
	IsReady() (ready bool)

	//Query
	SearchPostList(Keyword_Title string, Keyword_Author string,
		CreateAt_Start time.Time, CreateAt_End time.Time,
		Size int, Offset int, Sort []struct {
			SortBy      string
			IsAscending bool
		}) (totalcount int, result []ModelPost, err error)
	GetPost(UUID string) (result ModelPost, err error)
	AddPost(Source ModelPost) (err error)
	UpdatePost(Source ModelPost) (err error)
	DeletePost(UUID string, PasswordHash string) (err error)
}
