package mongoclient

import (
	"gopkg.in/mgo.v2"
	"log"
)

var mongoAddr = "mongodb://127.0.0.1:27017"
var dbName = "game"

var globalSession *mgo.Session

func init() {
	session, err := mgo.Dial(mongoAddr)
	if err != nil {
		log.Fatal()
	}
	session.SetMode(mgo.Monotonic, true)
	globalSession = session
}

func Insert(collection string, docs ...interface{}) error {
	if err := globalSession.Copy().DB(dbName).C(collection).Insert(docs...); err != nil {
		log.Println("insert failed ", err)
		return err
	}
	return nil
}

func FindOne(collection string, query interface{}, result interface{}) error {
	if err := globalSession.Copy().DB(dbName).C(collection).Find(query).One(result); err != nil {
		log.Println("find one err", err)
		return err
	}
	return nil
}

func FindAll(collection string, query interface{}, result interface{}) error {
	if err := globalSession.Copy().DB(dbName).C(collection).Find(query).All(result); err != nil {
		log.Println("find all err", err)
		return err
	}
	return nil
}
