package model

type Follower struct {
	ID  int  `json:"id"` // 粉丝id
	Follwed_id int  `json:"follwed_id"` // 关注的人id
}
