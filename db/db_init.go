package db

import (
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func InitDB() {
	dbobj := config.InitDBConnect()
	db, err := gorm.Open(dbobj.DBDriver, fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbobj.DBUser,
		dbobj.DBPasswd,
		dbobj.DBName))
	if err != nil {
		config.Debug(err)
		return
	}
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return dbobj.PREFIX + defaultTableName
	}

	db.SingularTable(true)

	db.Callback().Create().Register("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Register("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	fmt.Println("connect mysql success!")
	//err = db.Close()
	//fmt.Println(err.Error())
}

func GetOrm() *gorm.DB {
	return DB
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreateTime"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if updatedAtField, ok := scope.FieldByName("ModifiedOn"); ok {
			if updatedAtField.IsBlank {
				updatedAtField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}