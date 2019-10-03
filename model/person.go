package model

import (
	"database/sql"
	"errors"
	"time"
)

type Person struct {
	BaseModel
	Person_name     string  `gorm:"type:varchar(20);not null"`
	Id_card		string  `gorm:"type:varchar(20);not null"`
	//PersonInfoID   uint     `json:"person_infoid"`
	Password     string `gorm:"type:varchar(64);not null"`        //sha1
	Gender		 int `gorm:"type:tinyint(4);not null;default:0"`
	Login_ip      sql.NullString `gorm:"type:varchar(20)"`
	Login_time    time.Time  `gorm:"not null"`
	CreateTime     time.Time  `gorm:"not null" json:"create_time"`
	UpdateTime     time.Time  `gorm:"not null" json:"update_time"`
	Is_active     int      `gorm:"type:int;default:0"` //0禁用 1可用

}

func (p *Person) BeforeUpdate() (err error) {
	if (len(p.Id_card) != 18 ) {
		err = errors.New("user id is already greater than 1000")
	}
	return
}

//func GetPersonPage(pageNum int, pageSize int) ([]*Person, error) {
//	var persons []*Person
//	args1 := (pageNum-1) * pageSize
//	sqlstr := "select * from person limit ?,?;"
//	err := db.DB.Select(&persons,sqlstr,args1, pageSize)
//	if err != nil{
//		err_info :=fmt.Sprintf("Get person page failed, err:%s\n", err)
//		config.Error(err_info)
//		panic(err)
//	}
//	return persons, nil
//}

//func FindPersonByID(id int) (Person, error) {
//	var persons Person
//	sqlstr := "select * from person where id=?;"
//	err := db.DB.Get(&persons,sqlstr, id)
//	if err != nil{
//		err_info :=fmt.Sprintf("find the id: %d person failed, err:%s\n", id, err)
//		config.Error(err_info)
//		//panic(err)
//	}
//	return persons, nil
//}
//
//func AddtPersonInfo(p Person) (error) {
//	sqlstr := "insert into person(person_name,id_card,password,gender,login_ip) value(?,?,?,?,?);"
//	result , err := db.DB.Exec(sqlstr, p.Person_name,p.Id_card,p.Password,p.Gender,p.Login_ip)
//	if err != nil{
//		err_info :=fmt.Sprintf("add new person failed, err:%s\n", err)
//		config.Error(err_info)
//		panic(err)
//	}
//	affect, _ := result.RowsAffected()
//	if affect == 0{
//		err_info :=fmt.Sprintf("add new person failed, err:%s", err)
//		config.Error(err_info)
//		return errors.New(err_info)
//	}
//	return nil
//}


//func DelPersonByID(id int) error {
//	sqlstr := "delete from person where id=?;"
//	result, err := db.DB.Exec(sqlstr, id)
//	if err != nil{
//		err_info :=fmt.Sprintf("delete person failed, err:%s\n", err)
//		config.Error(err_info)
//		panic(err)
//	}
//	affect, _ := result.RowsAffected()
//	if affect == 0{
//		err_info :=fmt.Sprintf("The id to delete does not exist, plase enter again")
//		config.Error(err_info)
//		return errors.New(err_info)
//	}
//	return nil
//}
//
//func AcquirePassword(id uint) (string, error) {
//	sqlstr := "select password from person where id=?"
//	var oldPassword string
//	err := db.DB.Get(&oldPassword, sqlstr,id)
//	if err != nil{
//		err_info :=fmt.Sprintf("Unknow error happend err:%s\n", err)
//		config.Error(err_info)
//		panic(err)
//	}
//	return oldPassword, nil
//}
//
//func UpdatePassword(id uint, newPassword string) error {
//	sqlstr := "update person set password=? where id=?"
//	_, err := db.DB.Exec(sqlstr,newPassword, id)
//	if err != nil{
//		return errors.New("update password error")
//	}
//	return nil
//}