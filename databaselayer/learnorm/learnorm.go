package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type animal struct {
	gorm.Model
	Animaltype string `gorm:"type:TEXT"`
	Nickname   string `gorm:"type:TEXT"`
	Zone       int    `gorm:"type:INTEGER"`
	Age        int    `gorm:"type:INTEGER"`
}

func main() {
	db, err := gorm.Open("sqlite3", "dino.db")
	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println("connected")

	}
	defer db.Close()

	//
	//db.DropTableIfExists(&animal{})
	db.AutoMigrate(&animal{})
	//db.Table("dinos").CreateTable(&animal{})
	a := animal{
		Animaltype: "Trex2",
		Nickname:   "rex",
		Zone:       1,
		Age:        50,
	}
	db.Create(&a)
	//db.Table("dinos").Create
	a = animal{
		Animaltype: "Reptor",
		Nickname:   "Rap",
		Zone:       2,
		Age:        22,
	}
	db.Save(&a)

	//udpates
	//db.Table("animals").Where("nickname=? and zone=?", "Rap", 2).Update("age", 36)

	animals := []animal{}
	db.Find(&animals, "age >= ?", 35)
	fmt.Println(&animals)

}
