package model


type PersonInfo struct {
	ID           int
	PersonEmail  string `json:"person_email"`
	PersonAddr   string `json:"person_addr"`
	NickName   string `json:"nick_name"`
	Fans_count  int  `json:"fans_count"` // 粉丝数
	Follower_count int  `json:"follower_count"` // 关注数
}