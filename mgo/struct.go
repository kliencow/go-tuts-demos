package main

import (
	//"errors"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

var hostName string = "localhost"
var dbName string = "structFun"
var collectionName = "t"

var session *mgo.Session
var collection *mgo.Collection

// Open opens a connection to Mongo, sets up the db and the collection, and ensures appropriate indexes
func Open() {
	session, _ = mgo.Dial(hostName)
	session.SetMode(mgo.Monotonic, true)
	collection = session.DB(dbName).C(collectionName)
}

func Close() {
	session.Close()
}

type Storable interface {
	GetId() bson.ObjectId
	SetId(id bson.ObjectId)
	SetCreated()
	SetModified()
}

type Sub1 struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Common   string
	Unique1  string
	Created  time.Time
	Modified time.Time
}

func (s *Sub1) GetId() bson.ObjectId {
	return s.Id
}

func (s *Sub1) SetId(id bson.ObjectId) {
	s.Id = id
}

func (s *Sub1) SetCreated() {
	s.Created = time.Now()
}

func (s *Sub1) SetModified() {
	s.Modified = time.Now()
}

type Sub2 struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Common   string
	Unique2  string
	Created  time.Time
	Modified time.Time
}

func PrepGeneric(s Storable) {
	id := s.GetId()

	if !id.Valid() {
		id = bson.NewObjectId()
		s.SetCreated()
		s.SetId(id)
	}

	s.SetModified()
}

func main() {
	Open()
	defer Close()
	fmt.Println("Starting")

	s1 := new(Sub1)

	PrepGeneric(s1)
	id := s1.GetId()
	fmt.Println(id)
	_, err := collection.UpsertId(id, &s1)
	fmt.Println(err)
}
