package databaselayer

import (
	"database/sql"
	"fmt"
	"log"
)

//SQLHandler ...
type SQLHandler struct {
	*sql.DB
}

//GetAvailDynos ..
func (handler *SQLHandler) GetAvailDynos() ([]Animal, error) {
	result, err := handler.sendQuery("select * from Animals")

	return result, err
}

// GetDynosByNickname ...
func (handler *SQLHandler) GetDynosByNickname(n string) (Animal, error) {

	row := handler.QueryRow(fmt.Sprintf("select * from Animals where nickame='%s", n))
	a := Animal{}
	err := row.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
	return a, err
}

//GetDynosByType ...
func (handler *SQLHandler) GetDynosByType(t string) ([]Animal, error) {

	result, err := handler.sendQuery(fmt.Sprintf("select * from Animals where nickame='%s", t))

	return result, err
}

//AddAnimal ...
func (handler *SQLHandler) AddAnimal(a Animal) error {

	_, err := handler.Exec(fmt.Sprintf("Insert into Animals (animal_type, nickname, zone, age) values ('%s', '%s', '%d', '%d')", a.AnimalType, a.Nickname, a.Zone, a.Age))
	return err
}

// UpdateAnimal update the animal by it's nicknane
func (handler *SQLHandler) UpdateAnimal(a Animal, nname string) error {
	_, err := handler.Exec(fmt.Sprintf("Update  Animals set animal_type='%s', nickname='%s', zone='%d', age='%d' where nickname='%s'", a.AnimalType, a.Nickname, a.Zone, a.Age, nname))
	return err
}

func (handler *SQLHandler) sendQuery(q string) ([]Animal, error) {
	Animals := []Animal{}
	rows, err := handler.Query(q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		a := Animal{}
		err = rows.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
		if err != nil {
			log.Println(err)
			continue
		}
		Animals = append(Animals, a)
	}

	return Animals, rows.Err()

}
