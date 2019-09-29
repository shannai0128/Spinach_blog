package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/model"
	"regexp"
)

type PersonInfo struct {
	ID           int
	Person_email  sql.NullString
	Person_addr   sql.NullString
	Mobile       sql.NullString
	Nick_name   sql.NullString
}

func (p *PersonInfo) AddtPersonDetail() error {
	lenNick := len(p.Nick_name.String)
	if lenNick > 24 {
		err_msg := errors.New("please make sure nick name no more than 24 character")
		return err_msg
	}
	if len(p.Mobile.String) != 11{
		err_msg := errors.New("please input valid mobile number")
		return err_msg
	}
	reg, _ :=regexp.Compile("^([a-zA-Z0-9_-]+)@([a-zA-Z0-9_-]+)\\.com$")
	ok := reg.MatchString(p.Person_email.String)
	if !ok{
		err_msg := errors.New("please ensure valid email address")
		return err_msg
	}

	personInfo := model.PersonInfo{
		ID: p.ID,
		Person_email: p.Person_email,
		Person_addr: p.Person_addr,
		Mobile: p.Mobile,
		Nick_name: p.Nick_name,
	}
	err := model.AddtPersonDetail(personInfo)
	if err != nil{
		err_info :=fmt.Sprintf("add ID:%d user detail_info failed, err:%s", p.ID, err)
		config.Error(err_info)
		err_msg := errors.New(err_info)
		return err_msg
	}
	return nil
}

func (p *PersonInfo) EditPersonInfo() error {
	lenNick := len(p.Nick_name.String)
	if lenNick > 24 {
		err_msg := errors.New("please make sure nick name no more than 24 character")
		return err_msg
	}
	if len(p.Mobile.String) != 11{
		err_msg := errors.New("please input valid mobile number")
		return err_msg
	}
	reg, _ :=regexp.Compile("^([a-zA-Z0-9_-]+)@([a-zA-Z0-9_-]+)\\.com$")
	ok := reg.MatchString(p.Person_email.String)
	if !ok{
		err_msg := errors.New("please ensure valid email address")
		return err_msg
	}
	personInfo := model.PersonInfo{
		ID: p.ID,
		Person_email: p.Person_email,
		Person_addr: p.Person_addr,
		Mobile: p.Mobile,
		Nick_name: p.Nick_name,
	}
	err := model.EditPersonInfo(personInfo)
	if err != nil{
		err_info :=fmt.Sprintf("edit the person failed, err:%s", err)
		config.Error(err_info)
		err_msg := errors.New(err_info)
		return err_msg
	}
	return nil
}
