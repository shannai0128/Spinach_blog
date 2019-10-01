package model

import (
	"database/sql"
	"github.com/c479096292/Spinach_blog/db"
)

type PersonInfo struct {
	ID           int
	Person_email  sql.NullString
	Person_addr   sql.NullString
	Mobile       sql.NullString
	Nick_name   sql.NullString
	//Fans_count  int  `json:"fans_count"` // 粉丝数
	//Follower_count int  `json:"follower_count"` // 关注数
}


func AddtPersonDetail(p PersonInfo) (error) {
	sqlstr := "insert into person_info(id,Person_email,Person_addr,Mobile,Nick_name) value(?,?,?,?,?);"
	result , err := db.DB.Exec(sqlstr, p.ID,p.Person_email,p.Person_addr,p.Mobile,p.Nick_name)
	if err != nil {
		return err
	}
	affect, _ := result.RowsAffected()
	if affect == 0 {
		return err
	}
	return nil
}

func EditPersonInfo(p PersonInfo) error {
	sqlstr := "update person_info set Person_email=?,Person_addr=?,Mobile=?,Nick_name=? where id=?;"
	result , err := db.DB.Exec(sqlstr, p.Person_email, p.Person_addr,p.Mobile,p.Nick_name, p.ID)
	if err != nil{
		return err
	}
	affect, _ := result.RowsAffected()
	if affect == 0 {
		return err
	}
	return nil
}

func AcquireAllMobile() ([]sql.NullString, error) {
	var res []sql.NullString
	sqlstr := "select mobile from person_info;"
	err := db.DB.Select(&res, sqlstr)
	if err != nil{
		return nil,err
	}
	return res, nil
}
