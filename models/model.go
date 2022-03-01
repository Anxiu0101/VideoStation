package models

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"

	"VideoStation/pkg/setting"
)

var DB *gorm.DB

type Model struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt int64          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64          `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Setup Read the conf file and Open the database
func Setup() {
	var err error

	// pass conf to dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
	)
	fmt.Println(dsn)
	if dsn == ":@tcp()/?charset=utf8&parseTime=True&loc=Local" {
		dsn = "Anxiu:7391839@tcp(127.0.0.1:3306)/videostation?charset=utf8&parseTime=True&loc=Local"
	}
	fmt.Println(dsn)

	// open the database and buffer the conf
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.DatabaseSetting.TablePrefix, // set the prefix name of table
			SingularTable: true,                                // use singular table by default
		},
		Logger: logger.Default.LogMode(logger.Info), // set log mode
	})

	// some init set of database
	mysqlDB, err := DB.DB()
	if err != nil {
		log.Panicln("db.DB() err: ", err)
	}
	mysqlDB.SetMaxIdleConns(10)           // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	mysqlDB.SetMaxOpenConns(100)          // SetMaxOpenConns 设置打开数据库连接的最大数量
	mysqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime 设置了连接可复用的最大时间

	// set auto migrate
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{})
}

// CloseDB Close database
func CloseDB() {
	mysqlDB, _ := DB.DB()
	defer func(mysqlDB *sql.DB) {
		err := mysqlDB.Close()
		if err != nil {
			log.Panicln("DB.DB() err: ", err)
		}
	}(mysqlDB)
}
