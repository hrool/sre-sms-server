package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// sqllite
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Conn *gorm.DB

func init() {
	var err error
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "sms_" + defaultTableName
	}
	// Engine, err := gorm.Open("sqlite3", "admin/db.sqlite3")
	Conn, err = gorm.Open("postgres", "host=192.168.88.215 port=5432 user=postgres dbname=sre-sms password=mysecretpassword sslmode=disable")
	if err != nil {
		panic(err)
	}
	Conn.SingularTable(true)
}
