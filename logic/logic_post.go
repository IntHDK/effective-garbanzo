package logic

import (
	"effective-garbanzo/logic/common"
	"effective-garbanzo/logic/database"
	"errors"
	"time"
)

const (
	PARAM_SearchPostList_SortBy_CreateAt = "CreateAt"
)

func (l *LogicModule) SearchPostList(Keyword_Title string, Keyword_Author string,
	CreateAt_Start time.Time, CreateAt_End time.Time,
	Size int, Offset int, Sort []struct {
		SortBy      string
		IsAscending bool
	}) (totalcount int64, result []Post, err error) {
	if !l.database.IsReady() {
		totalcount = 0
		result = []Post{}
		err = errors.New(ERROR_DATASOURCE_NOTREADY)
		return
	}

	totalcount, src, err := l.database.SearchPostList(Keyword_Title, Keyword_Author, CreateAt_Start, CreateAt_End, Size, Offset, Sort)
	if err != nil {
		totalcount = 0
		result = []Post{}
		return
	}

	result = []Post{}
	for _, v := range src {
		result = append(result, Post{
			ID:       v.UUID,
			Author:   v.Author,
			Title:    v.Title,
			Context:  "",
			UpdateAt: v.UpdatedAt,
		})
	}

	return
}

func (l *LogicModule) GetPost(UUID string) (result Post, err error) {

	src, err := l.database.GetPost(UUID)
	if src.ID == 0 {
		result = Post{}
		err = errors.New(ERROR_DATASOURCE_ENTITYNOTFOUND)
		return
	}
	result = Post{
		ID:       src.UUID,
		Author:   src.Author,
		Title:    src.Title,
		Context:  src.Context,
		UpdateAt: src.UpdatedAt,
	}

	return
}

func (l *LogicModule) AddPost(Source Post, Password string) (UUID string, err error) {

	if err != nil {
		UUID = ""
		return
	}
	for {
		UUID = common.GenUUID()

		err = l.database.AddPost(database.ModelPost{
			UUID:    UUID,
			Title:   Source.Title,
			Context: Source.Context,
			Author:  Source.Author,
		}, Password)
		if err != database.ERROR_DUPLICATED {
			break
		}
	}

	return
}

func (l *LogicModule) UpdatePost(Source Post, Password string) (err error) {

	return
}

func (l *LogicModule) DeletePost(UUID string, Password string) (err error) {
	return
}
