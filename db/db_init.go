package db

import (
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var orm *gorm.DB

func init() {
	dbobj := config.InitDBConnect()
	var err error
	orm, err = gorm.Open(dbobj.DBDriver, fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbobj.DBUser,
		dbobj.DBPasswd,
		dbobj.DBName))
	if err != nil {
		config.Debug(err)
		return
	}
	orm.DB().SetMaxOpenConns(100)
	orm.DB().SetMaxIdleConns(10)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return dbobj.PREFIX + defaultTableName
	}

	orm.SingularTable(true)

	orm.Callback().Create().Register("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	orm.Callback().Update().Register("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	orm.DB().SetMaxIdleConns(10)
	orm.DB().SetMaxOpenConns(100)

	fmt.Println("connect mysql success!")
}

func GetOrm() *gorm.DB {
	return orm
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("CreateTime"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if updatedAtField, ok := scope.FieldByName("CreatedAt"); ok {
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