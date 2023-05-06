package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint   `gorm:"primarykey"`
	Code  string `gorm:"column:code"`
	Price uint   `gorm:"column:priceee"` // 自定义一下
	gorm.Model
}

type Video struct {
	AuthorId      uint   `gorm:"column:author_id"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	FavoriteCount int64  `gorm:"column:favourite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	Title         string `gorm:"column:title"`
	gorm.Model
}

type User struct {
	Name           string  `json:"name"`
	Password       string  `json:"password"`
	FavoriteVideos []Video `gorm:"many2many:favorite" json:"favorite_videos"`
	gorm.Model
}

func main() {
	// 建立数据库连接，建议使用viper读取配置
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:lht12138@tcp(1.15.78.83:9887)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}

	var user User
	err = db.First(&user, "123").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(err)
	}

	//// 建表，若已存在不会重复创建
	//err = db.AutoMigrate(&Product{})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	////创建一条
	//p := &Product{Code: "A18"}
	//res := db.Create(p)
	//fmt.Println(res.Error)
	//fmt.Println(p.ID)
	//
	////	创建多条
	//products := []*Product{{Code: "A11"}, {Code: "A12"}, {Code: "A13"}}
	//res = db.Create(products)
	//fmt.Println(res.Error)
	//for _, p := range products {
	//	fmt.Println(p.ID)
	//}
}
