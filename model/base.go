package model

import (
	"time"
)

type BaseModel struct {
	CreateTime time.Time `json:"create_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

