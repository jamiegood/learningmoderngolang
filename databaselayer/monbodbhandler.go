package databaselayer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongodbHandler ...
type MongodbHandler struct {
	*mgo.Session
}

//NewMongodbHandler ...
func NewMongodbHandler(connection string) (*MongodbHandler, error) {

	s, err := mgo.Dial(connection)
	return &MongodbHandler{
		Session: s,
	}, err
}

//GetAvailDynos ...
func (handler *MongodbHandler) GetAvailDynos() ([]Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	animals := []Animal{}
	err := s.DB("Dino").C("animals").Find(nil).All(&animals)
	return animals, err
}

//GetDynosByNickname ...
func (handler *MongodbHandler) GetDynosByNickname(nickname string) (Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	animal := Animal{}
	err := s.DB("Dino").C("animals").Find(bson.M{"nickname": nickname}).One(&animal)
	return animal, err
}

//GetDynosByType ...
func (handler *MongodbHandler) GetDynosByType(t string) ([]Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	animals := []Animal{}
	err := s.DB("Dino").C("animals").Find(bson.M{"animal_type": t}).All(&animals)

	return animals, err
}

//AddAnimal ...
func (handler *MongodbHandler) AddAnimal(animal Animal) error {
	s := handler.getFreshSession()
	defer s.Close()

	return s.DB("Dino").C("animals").Insert(animal)
}

//UpdateAnimal ...
func (handler *MongodbHandler) UpdateAnimal(a Animal, nickname string) error {
	s := handler.getFreshSession()
	defer s.Close()

	return s.DB("Dino").C("animals").Update(bson.M{"nickname": nickname}, a)
}

//getFreshSession ...
func (handler *MongodbHandler) getFreshSession() *mgo.Session {
	return handler.Session.Copy()
}
