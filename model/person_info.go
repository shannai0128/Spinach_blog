package model

import (
	"database/sql"
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/db"
)

type PersonInfo struct {
	ID           int
	Person_email  sql.NullString `json:"person_email"`
	Person_addr   sql.NullString `json:"person_addr"`
	Mobile       sql.NullString `json:"mobile"`
	Nick_name   sql.NullString `json:"nick_name"`
	//Fans_count  int  `json:"fans_count"` // 粉丝数
	//Follower_count int  `json:"follower_count"` // 关注数
}



func AddtPersonDetail(p Person) (error) {
	sqlstr := "insert into person_info(person_name,id_card,password,gender,login_ip) value(?,?,?,?,?);"
	_ , err := db.DB.Exec(sqlstr, p.Person_name,p.Id_card,p.Password,p.Gender,p.Login_ip)
	if err != nil{
		err_info :=fmt.Sprintf("create new article failed, err:%s\n", err)
		config.Error(err_info)
		return err
	}
	return nil
}

//func EditPersonInfo(p Person) error {
//	sqlstr := "update person set category_id=?,content=?,title=?,view_count=?,person_name=?,summary=?,origin=? where id=?;"
//	_ , err := db.DB.Exec(sqlstr, a.Category_id, a.Content, a.Title,a.View_count, a.Person_name, a.Summary, a.Origin, a.ID)
//	if err != nil{
//		err_info :=fmt.Sprintf("edit the article failed, err:%s\n", err)
//		config.Error(err_info)
//		return err
//	}
//	return nil
//}
