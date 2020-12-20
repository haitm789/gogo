package database

import (
	"fmt"
	"os"

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
	dc, err := loadConfig()
	if err != nil {
		panic(err.Error())
	}
	DBMS := dc.DbMs
	USER := dc.User
	PASS := dc.Pass
	PROTOCOL := fmt.Sprintf("tcp(%s:%s)", dc.Host, dc.Port)
	DBNAME := dc.DbName
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true&loc=Local"
	fmt.Println("db connect:", CONNECT)
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	return db
}

type config struct {
	DbMs   string
	User   string
	Pass   string
	DbName string
	Host   string
	Port   string
}

func loadConfig() (dc config, err error) {
	dc.DbMs = "mysql"
	dc.DbName = os.Getenv("DB_NAME")
	dc.User = os.Getenv("DB_USER")
	dc.Pass = os.Getenv("DB_PASSWORD")
	dc.Host = os.Getenv("DB_HOST")
	dc.Port = os.Getenv("DB_PORT")
	return dc, nil
}
