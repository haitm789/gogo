package database

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type mysql struct {
	Conn *gorm.DB
}

func NewMySql() *mysql {
	c := connect()
	return &mysql{
		Conn: c,
	}
}

func (m *mysql) Connection() *gorm.DB {
	return m.Conn
}

func connect() *gorm.DB {
	dialect, conn := conn()

	db, err := gorm.Open(dialect, conn)
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	return db
}

func conn() (string, string) {
	fmt.Println("viper database:", viper.GetString("database.host"))

	host := "gogo-mysql"
	port := "33061"
	username := "pandog"
	password := "pandog"
	database := "gogo"
	dialect := "mysql"

	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)

	return dialect, conn
}
