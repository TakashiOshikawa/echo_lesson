package dao

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

// localtest mysql
var Db, err = gorm.Open("mysql", "root:password@tcp(192.168.99.100:3306)/quill_lesson?charset=utf8&parseTime=True")
