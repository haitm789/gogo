package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Mysql struct {
	Conn *gorm.DB
}

func NewMySql() *Mysql {
	return &Mysql{
		Conn: connect(),
	}
}

func (m *Mysql) Connection() *gorm.DB {
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
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	dialect := os.Getenv("DB_DIALECT")

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
