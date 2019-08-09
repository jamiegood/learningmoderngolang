package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type animal struct {
	Animaltype string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1")

	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println("connected")

	}
	defer session.Close()

	// get collectoin
	animalcollection := session.DB("Dino").C("animals")

	// a := animal{
	// 	Animaltype: "big",
	// 	Nickname:   "biggie",
	// 	Zone:       3,
	// 	Age:        5,
	// }

	// animals := []interface{}{
	// 	animal{
	// 		Animaltype: "biga",
	// 		Nickname:   "biggiea",
	// 		Zone:       3,
	// 		Age:        5,
	// 	},
	// 	animal{
	// 		Animaltype: "bigB",
	// 		Nickname:   "biggieB",
	// 		Zone:       3,
	// 		Age:        5,
	// 	},
	// 	animal{
	// 		Animaltype: "bigC",
	// 		Nickname:   "biggieC",
	// 		Zone:       3,
	// 		Age:        5,
	// 	},
	// }

	//err = animalcollection.Insert(animals...)
	if err != nil {
		fmt.Println(err)
	}

	// err = animalcollection.Update(bson.M{"nickname": "biggieC"}, bson.M{"$set": bson.M{"age": 99}})
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = animalcollection.Remove(bson.M{"nickname": "biggieB"})
	// if err != nil {
	// 	fmt.Println(err)
	// }

	query := bson.M{
		"age": bson.M{
			"$gt": 50,
		},
		"zone": bson.M{
			"$in": []int{1, 2, 3},
		},
	}

	//create array of animals
	results := []animal{}
	err = animalcollection.Find(query).One(&results)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results)

}
