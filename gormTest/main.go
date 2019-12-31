package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Age          int
	Name         string  `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
	Num          int     `gorm:"AUTO_INCREMENT"` // 自增
	Test         int
	Test1         int
}

type Product struct {
	gorm.Model
	Code string
	Price uint
}


func main() {
	db, err := gorm.Open("mysql", "root:vtesoho1@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	//db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{},&User{})

	//db.Create(&Product{Code: "L1212", Price: 1000})
	db.Create(&User{Age: 1, Name: "vs",Test:2})
	//db.Create(&User{Code: "L1212", Price: 1000})


	//a := db.HasTable(&User{})
	//b := db.CreateTable(&User{})
	//println("asdasdasdasd",a,"-------------",b)
	time.Sleep(time.Second * 3)
	println("end")
	//db.Close()
}
