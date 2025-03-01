package service

import (
	"ztalk/internal/models"
	"ztalk/internal/repository/mysql"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommunityById(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailById(id)
}
