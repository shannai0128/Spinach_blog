package model

type Article struct {
	ID             uint       `json:"aid"`
	//CreatedAt      time.Time  `json:"-"`
	//UpdatedAt      time.Time  `json:"-"`
	BaseModel
	//DeletedAt      *time.Time `sql:"index" json:"-"`
	Person_id         uint     `json:"Person_id"`
	CategoryID 	   uint       `json:"cid"`
	Title          string     `json:"title"`
	Summary		   string	   `json:"summary"`
	Content        string     `json:"content"`
	View_count     int        `json:"view_count"`
	//img          int        `json:"img"`
	Origin         int        `json:"origin"` //是否原创 1原创 0转载
	State          int        `json:"state"`     //0正常发布 2并未发布(草稿箱)
	Violat_reason  string     `json:"violat_reason"`
	praise         int		  `json:"praise"`
	Comment_id       string     `json:"comment_id"`
	Comment_count  int    	  `json:"comment_count"`
	Praise         int		  `json:"praise"`
}

