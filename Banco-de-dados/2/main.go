package main

import (
	"gorm.io/driver/mysql"
)
import "gorm.io/gorm"

type Product struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//db.Create(&Product{
	//	Name:  "Notebook",
	//	Price: 1000.00,
	//})

	//batch
	//	products := []Product{
	//		{Name: "Notebook", Price: 1000.00},
	//		{Name: "Mouse", Price: 50.00},
	//		{Name: "Keyboard", Price: 100.00},
	//	}
	//		db.Create(&products)

	//select
	//var product Product
	//db.First(&product, 1)
	//fmt.Println(product)
	//db.First(&product, "name =?", "Mouse")
	//fmt.Println(product)

	//select all
	//var products []Product
	//db.Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	//offset
	//var products []Product
	//db.Limit(2).Offset(2).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	//where
	//var product []Product
	//db.Where("price > ?", 100).Find(&product)
	//db.Where("name LIKE ?", "%k%").Find(&product)
	//for _, product := range product {
	//	fmt.Println(product)
	//}

	//update
	//var p Product
	//db.First(&p, 1)
	//p.Name = "New Mouse"
	//db.Save(&p)

	//delete
	//var p2 Product
	//db.First(&p2, 1)
	//fmt.Println(p2)
	//db.Delete(&p2)

}
