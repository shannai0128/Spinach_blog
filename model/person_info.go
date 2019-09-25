package model

import "database/sql"

type PersonInfo struct {
	ID           int
	PersonEmail  sql.NullString `json:"person_email"`
	PersonAddr   sql.NullString `json:"person_addr"`
	NickName   sql.NullString `json:"nick_name"`
	//Fans_count  int  `json:"fans_count"` // 粉丝数
	//Follower_count int  `json:"follower_count"` // 关注数
}