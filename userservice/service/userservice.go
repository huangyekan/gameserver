package service

import (
	"gmserver/userservice/mongoclient"
	"gopkg.in/mgo.v2/bson"
	"log"
	"gmserver/userservice/model"
)

type UserService int

func (u *UserService) GetUserByAccount(args map[string]interface{}, reply *model.User) error {
	err := mongoclient.FindOne("user", bson.M{"Account" : args["Account"], "valid" : true}, &reply)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}