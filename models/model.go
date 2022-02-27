package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"

	"VideoStation/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt int64          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64          `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Setup Read the config file and Open the database
func Setup() {
	var (
		err                                       error
		dbName, user, password, host, tablePrefix string
	)

	// Read value from config attribute name
	dbName = setting.DatabaseSetting.Name
	user = setting.DatabaseSetting.User
	password = setting.DatabaseSetting.Password
	host = setting.DatabaseSetting.Host
	tablePrefix = setting.DatabaseSetting.TablePrefix

	// pass config to dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)

	// open the database and buffer the config
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // set the prefix name of table
			SingularTable: true,        // use singular table by default
		},
		Logger: logger.Default.LogMode(logger.Info), // set log mode
	})

	mysqlDB, err := db.DB()
	if err != nil {
		log.Panicln("db.DB() err: ", err)
	}

	// some init set of database
	mysqlDB.SetMaxIdleConns(10)  // set max idle connections
	mysqlDB.SetMaxOpenConns(100) // set open connections, default is 0 (unlimited)
}

// CloseDB Close database
func CloseDB() {
	mysqlDB, _ := db.DB()
	defer mysqlDB.Close()
}
