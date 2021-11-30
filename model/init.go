package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
	"github.com/spf13/viper"

	// 导入mysql驱动 下划线表示只调用mysql包的init函数
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	LocalDB *gorm.DB // 本地连接
	CloudDB *gorm.DB // 云服务器连接
}

var DB *Database

// 初始化2个连接
func (db *Database) Init() {
	DB = &Database{
		LocalDB: GetLocalDB(),
		CloudDB: GetCloudDB(),
	}
}

func GetLocalDB() *gorm.DB {
	// 读取对应连接中的配置
	return openDB(viper.GetString("local_db.username"),
		viper.GetString("local_db.password"),
		viper.GetString("local_db.addr"),
		viper.GetString("local_db.name"))
}

func GetCloudDB() *gorm.DB {
	return openDB(viper.GetString("cloud_db.username"),
		viper.GetString("cloud_db.password"),
		viper.GetString("cloud_db.addr"),
		viper.GetString("cloud_db.name"))
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxOpenConns(20000) // 设置最大打开的连接数，默认值为0表示不限制
	db.DB().SetMaxIdleConns(0)     // 设置闲置的连接数，设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用
}

// 建立数据库连接
func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	// 建立数据库连接
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setupDB(db)

	return db
}

func (db *Database) Close() {
	DB.LocalDB.Close()
	DB.CloudDB.Close()
}
