package mysql

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"ztalk/internal/models"
	"ztalk/pkg/message"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)
	if errors.Is(err, sql.ErrNoRows) {
		zap.L().Warn("there is no community in database")
		err = nil
	}
	return
}

func GetCommunityDetailById(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time from community where community_id = ?`
	err = db.Get(community, sqlStr, id)
	if errors.Is(err, sql.ErrNoRows) {
		err = message.ErrInvalidID
	}
	return
}
