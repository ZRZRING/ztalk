package models

import "time"

type Community struct {
	Id   int64  `db:"community_id" json:"id"`
	Name string `db:"community_name" json:"name"`
}

type CommunityDetail struct {
	Id           int64     `db:"community_id" json:"id"`
	Name         string    `db:"community_name" json:"name"`
	Introduction string    `db:"introduction" json:"introduction,omitempty"`
	CreateTime   time.Time `db:"create_time" json:"create_time"`
}
