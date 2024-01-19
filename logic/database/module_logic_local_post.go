package database

import (
	"effective-garbanzo/logic/common"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func (module DatabaseModule_Local) SearchPostList(
	Keyword_Title string, Keyword_Author string,
	CreateAt_Start time.Time, CreateAt_End time.Time,
	Size int, Offset int, Sort []struct {
		SortBy      string
		IsAscending bool
	}) (totalcount int64, result []ModelPostListRecord, err error) {
	totalcount = 0
	result = []ModelPostListRecord{}
	err = nil

	var qtx *gorm.DB = module.db

	if Keyword_Title != "" {
		qtx = qtx.Where(CONDITION_FIELD_POST_TITLE+" LIKE ?", "%"+Keyword_Title+"%")
	}
	if Keyword_Author != "" {
		qtx = qtx.Where(CONDITION_FIELD_POST_AUTHOR+" LIKE ?", "%"+Keyword_Author+"%")
	}
	if !CreateAt_Start.IsZero() {
		qtx = qtx.Where(CONDITION_FIELD_POST_UPDATEDAT+" BETWEEN ? AND ?", CreateAt_Start, CreateAt_End)
	}

	orderstrs := []string{}
	for _, v := range Sort {
		ascstr := "asc"
		if !v.IsAscending {
			ascstr = "desc"
		}
		orderstr := fmt.Sprintf("%s %s", v.SortBy, ascstr)
		orderstrs = append(orderstrs, orderstr)
	}

	for _, orderstr := range orderstrs {
		qtx = qtx.Order(orderstr)
	}

	if Size > 0 {
		qtx = qtx.Limit(Size).Offset(Offset)
	}

	qresult := qtx.Model(&ModelPost{}).Find(&result)
	if qresult.Error != nil {
		err = qresult.Error
		result = []ModelPostListRecord{}
		totalcount = 0
		//TODO: 에러 처리
		return
	}
	totalcount = qresult.RowsAffected
	err = nil

	return
}

func (module DatabaseModule_Local) AddPost(Source ModelPost, Password string) (err error) {
	err = nil

	pwhash, err := common.PasswordHash(Password)
	if err != nil {
		return
	}

	result := module.db.Create(&ModelPost{
		UUID:         Source.UUID,
		Title:        Source.Title,
		Context:      Source.Context,
		Author:       Source.Author,
		PasswordHash: pwhash,
	})
	if result.Error != nil {
		err = result.Error
	}

	return
}

func (module DatabaseModule_Local) GetPost(UUID string) (result ModelPost, err error) {
	result = ModelPost{}
	err = nil

	var src ModelPost
	dbres := module.db.Where(&ModelPost{UUID: UUID}).First(&src)
	if dbres.Error != nil {
		if dbres.Error == gorm.ErrRecordNotFound {
			return
		} else {
			err = dbres.Error
			return
		}
	}
	result = src

	return
}

func (module DatabaseModule_Local) UpdatePost(Source ModelPost, Password string) (err error) {
	err = module.db.Transaction(func(tx *gorm.DB) error {
		var src ModelPost
		dbres := tx.Where(&ModelPost{UUID: Source.UUID}).First(&src)
		if dbres.Error != nil {
			return dbres.Error
		}

		if !common.ComparePasswordHash(src.PasswordHash, Password) {
			return errors.New(ERROR_PASSWORDHASH_INCORRECT)
		}

		dbres = tx.Model(&src).Updates(map[string]interface{}{
			"Title":   Source.Title,
			"Context": Source.Context,
			"Author":  Source.Author,
		})
		if dbres.Error != nil {
			return dbres.Error
		}

		return nil
	})

	return
}

func (module DatabaseModule_Local) DeletePost(UUID string, Password string) (err error) {
	err = module.db.Transaction(func(tx *gorm.DB) error {
		var src ModelPost
		dbres := tx.Where(&ModelPost{UUID: UUID}).First(&src)
		if dbres.Error != nil {
			return dbres.Error
		}

		if !common.ComparePasswordHash(src.PasswordHash, Password) {
			return errors.New(ERROR_PASSWORDHASH_INCORRECT)
		}

		dbres = tx.Delete(&src)
		if dbres.Error != nil {
			return dbres.Error
		}

		return nil
	})

	return
}
