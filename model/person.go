package model

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/db"
	"github.com/jinzhu/gorm"
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
	Login_time    time.Time
	CreateTime     time.Time
	UpdateTime     time.Time
	Is_active     int      `gorm:"type:int;default:0"` //0禁用 1可用

}

func (p *Person) db() *gorm.DB {
	return db.GetOrm()
}

func (p *Person) BeforeUpdate() (err error) {
	if (len(p.Id_card) != 18 ) {
		err = errors.New("id card greater than 18 character")
	}
	return
}

func (p *Person) GetTableTotal() int {
	var v int
	err := p.db().Model(&Article{}).Count(&v).Error
	if err !=nil {
		err_info :=fmt.Sprintf("Get article Total failed, err:%v\n", err)
		config.Error(err_info)
	}
	fmt.Println(v)
	return v
}

func (p *Person) GetArticlesPage(pageNum int, pageSize int) ([]*Person, error) {
	var persons []*Person
	args1 := (pageNum-1) * pageSize
	err := p.db().Limit(pageSize).Offset(args1).Find(&persons).Error
	if err != nil {
		err_info :=fmt.Sprintf("Get person page failed, err:%s\n", err.Error())
		config.Error(err_info)
		panic(err_info)
	}
	return persons, nil
}

func (p *Person) FindPersonByID(id int) (Person, error) {
	var persons Person
	err := p.db().Where("id = ?",id).First(&persons).Error
	if err != nil{
		err_info :=fmt.Sprintf("find the id: %d person failed, err:%s\n", id, err)
		config.Error(err_info)
		//panic(err)
	}
	return persons, nil
}

func (p *Person) AddtPersonInfo() (error) {

	err := p.db().Create(&p).Error
	if err != nil{
		err_info :=fmt.Sprintf("create new person failed, err:%s\n",err)
		config.Error(err_info)
		return err
	}
	return nil
}


func (p *Person) DelPersonByID() error {
	err := p.db().Delete(&p).Error
	if err != nil{
		err_info :=fmt.Sprintf("create person failed, err:%s\n",err)
		config.Error(err_info)
		return err
	}
	return nil
}

func (p *Person) AcquirePassword(id uint) (string, error) {
	var personObj Person
	err := p.db().Where("id = ?", id).Select("person_name").Find(&personObj).Error
	if err != nil{
		err_info :=fmt.Sprintf("Unknow error happend err:%s\n", err)
		config.Error(err_info)
		return "", err
	}
	return personObj.Password, nil
}

func (p *Person) UpdatePassword() error {
	err := p.db().Model(&p).Where("id = ?",p.ID).Updates(&p).Error
	if err != nil{
		return errors.New("update password error")
	}
	return nil
}