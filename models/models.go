package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"qianshen/pkg/setting"
)

var (
	db *gorm.DB
)

func DB() *gorm.DB {
	return db
}

func migrate() {
	models := []interface{}{User{}}
	db.AutoMigrate(models...)
}

// Setup initializes the database instance
func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}
	if setting.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	//migrate()// 不让他自动创建表，因为float不可以为无符号型的数据
	//db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	go migrate()
}
