package database

import (
	"fmt"

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
	// dialect := os.Getenv("DB_DIALECT")
	// username := os.Getenv("DB_USERNAME")
	// password := os.Getenv("DB_PASSWORD")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// database := os.Getenv("DB_DATABASE")

	host := "gogo-mysql"
	port := "3306"
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
