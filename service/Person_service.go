package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/model"
	"github.com/c479096292/Spinach_blog/utils"
	"time"
)

type Person struct {
	ID            int
	Name 		  string
	PersonName     string
	IdCard		string
	PersonInfoID   uint
	PassWord     string
	Gender		 int
	article_id   int
	LoginIP      sql.NullString
	LoginTime    time.Time
	IsActive     int
}

func (p *Person) Count() (int) {
	person := model.Person{}
	return person.GetTableTotal()
}

func (p *Person) GetPersonPage(pageNum int, pageSize int) ([]*model.Person) {
	person := model.Person{}
	persons, err := person.GetArticlesPage(pageNum, pageSize)
	if err != nil{
		if err != nil{
			config.Error(fmt.Sprintf("query person paged error: %s\n",err))
		}
	}
	return persons
}

func (p *Person) FindPersonByID(id int) (model.Person) {
	person := model.Person{}
	persons, err := person.FindPersonByID(id)
	if err != nil{
		if err != nil{
			config.Error(fmt.Sprintf("search id:%d person error: %s\n",id, err))
		}
	}
	return persons
}

func (p *Person) AddtNewPerson() error {
	pwd := len(p.PassWord)
	if pwd < 8 {
		err_info := errors.New("please check password, make sure password length is greater than 8")
		return err_info
	}
	idCard := len(p.IdCard)
	if idCard != 18 {
		err_info := errors.New("please check and input valid idCard")
		return err_info
	}
	p.PassWord = utils.EncodeMD5(p.PassWord)

	person := model.Person{
		Person_name:p.PersonName,
		Id_card:p.IdCard,
		Password:p.PassWord,
		Login_ip:p.LoginIP,
		Gender:p.Gender,
		Login_time:time.Now(),
		//CreateTime:time.Now(),
		UpdateTime:time.Now(),
	}
	err := person.AddtPersonInfo()
	if err !=nil {
		config.Error(fmt.Sprintf("add person error: %s\n",err))
		return err
	}
	return nil
}

func (p *Person) DelPersonByID(id uint) error {
	person := model.Person{}
	person.ID = id
	err := person.DelPersonByID()
	if err !=nil {
		config.Error(fmt.Sprintf("delete person error: %s\n",err))
		return err
	}
	return nil
}
//
func (p *Person) ModifyPassword(id uint, oldPwd, newPwd string) error {
	person := model.Person{}
	acqPassword, err := person.AcquirePassword(id)
	oldPassword := utils.EncodeMD5(oldPwd)
	if oldPassword != acqPassword {
		return errors.New("password error, please try again")
	}
	newPassword := utils.EncodeMD5(newPwd)
	person.Password = newPassword
	err = person.UpdatePassword()
	if err != nil{
		err_info :=fmt.Sprintf("Unknow error happend err:%s\n", err)
		return errors.New(err_info)
		//panic(err)
	}
	return nil
}