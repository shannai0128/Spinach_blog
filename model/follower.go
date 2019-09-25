package model

type Follower struct {
	ID  int  `json:"id"`
	//User_id int `json:"user_id"`
	// 没有设置sql.null*  操作此表不要涉及*
	Followed_id []int  `json:"follwed_id"` // 关注的人id
	Followed []int `json:"followed"` // 粉丝
	Status int `json:"status"` // 0 关注别人,1 被别人关注, 2 互相关注
}
