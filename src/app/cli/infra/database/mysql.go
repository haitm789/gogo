package database

import (
	"fmt"
	"os"

	// "github.com/spf13/viper"

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
	fmt.Println("G" + os.Getenv("DB_USERNAME"))
	dialect, conn := conn()

	db, err := gorm.Open(dialect, conn)
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	return db
}

func conn() (string, string) {
	host := "gogo-mysql"
	port := "3306"
	username := "pandog"
	password := "pandog"
	database := "gogo"
	dialect := "mysql"

	// 	DB_HOST="gogo-mysql"
	// DB_PORT=3306
	// DB_USERNAME=pandog
	// DB_PASSWORD=pandog
	// DB_DATABASE=gogo
	// DB_DIALECT=mysql

	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// username := os.Getenv("DB_USERNAME")
	// password := os.Getenv("DB_PASSWORD")
	// database := os.Getenv("DB_DATABASE")
	// dialect := os.Getenv("DB_DIALECT")

	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)
	fmt.Println("vvvv")
	fmt.Println(os.Getenv("DB_HOST"))

	return dialect, conn
}
