package service

import (
	"database/sql"
	"errors"
	"ztalk/internal/models"
	"ztalk/internal/mysql"
	"ztalk/pkg/message"

	"go.uber.org/zap"
)

func GetCommunityList() (data []*models.Community, err error) {
	data, err = mysql.GetAllCommunities()
	if errors.Is(err, sql.ErrNoRows) {
		zap.L().Warn("there is no community in database")
		err = nil
	}
	return
}

func GetCommunityByID(id int64) (data *models.Community, err error) {
	data, err = mysql.GetCommunityByID(id)
	if errors.Is(err, sql.ErrNoRows) {
		err = message.ErrInvalidID
		return
	}
	return
}
