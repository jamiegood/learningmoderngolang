package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func main() {
	db, err := sql.Open("mysql", "root:@/Dino")
	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println("connected")

	}
	defer db.Close()

	rows, err := db.Query("select * from Dino.animals where age > ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{}

	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}

		animals = append(animals, a)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(animals)

	//return a single method
	row := db.QueryRow("select * from Dino.animals where id=?", 1)
	a := animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
	if err != nil {

	}
	fmt.Println(a)

	//insert a row
	result, err := db.Exec("insert into Dino.animals (animal_type, nickname, zone, age) values ('Titanosours', 'titan', 3, 22)")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

}
