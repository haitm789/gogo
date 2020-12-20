package adapter

import (
	"github.com/jinzhu/gorm"
)

type Database interface {
	Connection() *gorm.DB
}
