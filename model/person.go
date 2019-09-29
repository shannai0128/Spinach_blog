package model

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/db"
)

type Person struct {
	ID          uint

	BaseModel
	Person_name     string
	Id_card		string
	//PersonInfoID   uint     `json:"person_infoid"`
	Password     string `json:"password"`        //sha1
	Gender		 int
	Login_ip      sql.NullString
	Login_time    string
	Is_active     int      `json:"is_active"` //0可用 1禁用

}

func GetPersonPage(pageNum int, pageSize int) ([]*Person, error) {
	var persons []*Person
	args1 := (pageNum-1) * pageSize
	sqlstr := "select * from person limit ?,?;"
	err := db.DB.Select(&persons,sqlstr,args1, pageSize)
	if err != nil{
		err_info :=fmt.Sprintf("Get person page failed, err:%s\n", err)
		config.Error(err_info)
		panic(err)
	}
	return persons, nil
}

func FindPersonByID(id int) (Person, error) {
	var persons Person
	sqlstr := "select * from person where id=?;"
	err := db.DB.Get(&persons,sqlstr, id)
	if err != nil{
		err_info :=fmt.Sprintf("find the id: %d person failed, err:%s\n", id, err)
		config.Error(err_info)
		//panic(err)
	}
	return persons, nil
}

func AddtPersonInfo(p Person) (error) {
	sqlstr := "insert into person(person_name,id_card,password,gender,login_ip) value(?,?,?,?,?);"
	result , err := db.DB.Exec(sqlstr, p.Person_name,p.Id_card,p.Password,p.Gender,p.Login_ip)
	if err != nil{
		err_info :=fmt.Sprintf("add new person failed, err:%s\n", err)
		config.Error(err_info)
		panic(err)
	}
	affect, _ := result.RowsAffected()
	if affect == 0{
		err_info :=fmt.Sprintf("add new person failed, err:%s", err)
		config.Error(err_info)
		return errors.New(err_info)
	}
	return nil
}


func DelPersonByID(id int) error {
	sqlstr := "delete from person where id=?;"
	result, err := db.DB.Exec(sqlstr, id)
	if err != nil{
		err_info :=fmt.Sprintf("delete person failed, err:%s\n", err)
		config.Error(err_info)
		panic(err)
	}
	affect, _ := result.RowsAffected()
	if affect == 0{
		err_info :=fmt.Sprintf("The id to delete does not exist, plase enter again")
		config.Error(err_info)
		return errors.New(err_info)
	}
	return nil
}

func AcquirePassword(id uint) (string, error) {
	sqlstr := "select password from person where id=?"
	var oldPassword string
	err := db.DB.Get(&oldPassword, sqlstr,id)
	if err != nil{
		err_info :=fmt.Sprintf("Unknow error happend err:%s\n", err)
		config.Error(err_info)
		panic(err)
	}
	return oldPassword, nil
}

func UpdatePassword(id uint, newPassword string) error {
	sqlstr := "update person set password=? where id=?"
	_, err := db.DB.Exec(sqlstr,newPassword, id)
	if err != nil{
		return errors.New("update password error")
	}
	return nil
}