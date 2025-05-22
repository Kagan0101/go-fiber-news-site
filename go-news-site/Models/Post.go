package Models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Info struct {
	gorm.Model
	Title,Information,Picture_url,Categories string
	Information_id,Categories_id int
}



func (info Info) Migrate(){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&info)
}


func (info Info) Add(){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}

	db.Create(&info)
}

func (info Info) Get(where ...interface{}) Info {
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return info
	}
	db.First(&info,where...)
	return info
}

func (info Info) GetAll(where ...interface{}) []Info {
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var posts []Info
	db.Find(&posts,where...)
	return posts
}

func (info Info) Update(column string,value interface{})  {
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&dsn).Update(column,value)
}

func (info Info) Updates(data Info){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&info).Updates(data)
}

func (info Info) Delete()  {
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&info,info.ID)
}