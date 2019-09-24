package app

import (
	"github.com/c479096292/Spinach_blog/common"
	"github.com/jmoiron/sqlx"
)

type App struct {
	DB *sqlx.DB
	common.Response
}





