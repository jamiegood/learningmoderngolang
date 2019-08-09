package databaselayer

import "errors"

// constants, this is how you do an enum in go
const (
	MYSQL uint8 = iota
	SQLITE
	POSTGRESQL
	MONGODB
)

// DinoDBHandler comment here
type DinoDBHandler interface {
	GetAvailDynos() ([]Animal, error)
	GetDynosByNickname(string) (Animal, error)
	GetDynosByType(string) ([]Animal, error)
	AddAnimal(Animal) error
	UpdateAnimal(Animal, string) error
}

//Animal comment
type Animal struct {
	ID         int    `bson:"-"`
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

//ErrDBTypeNotSupported comment
var ErrDBTypeNotSupported = errors.New("the Database type provided is not supported")

// GetDatabaseHandler factory
func GetDatabaseHandler(dbtype uint8, connection string) (DinoDBHandler, error) {

	switch dbtype {
	case MYSQL:
		return NewMySQLHandler(connection)
	case MONGODB:
		return NewMongoDBHandler(connection)
	case SQLITE:
		return NewSQLiteHandler(connection)
	case POSTGRESQL:
		return NewPQHandler(connection)

	}
	return nil, ErrDBTypeNotSupported
}
