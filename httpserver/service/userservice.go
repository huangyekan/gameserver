package service

import (
	"github.com/huangyekan/gameserver/userservice/model"
	"log"
	"gmserver/httpserver/util/md5"
)

type UserService int

var tokenMapping map[string]string = make(map[string]string)

func (u *UserService) IsValidUser(account string, password string) bool {
	params := map[string]interface{}{"account": account}
	var user model.User
	err := RemoteService("UserService", "GetUserByAccounts", params, &user)
	if err != nil {
		log.Println(err)
		return false
	}

	if user.Password == md5.Encode(password) {
		return true
	}
	return false
}

func (u *UserService) CheckToken(token string) bool {
	if _, ok := tokenMapping[token]; ok {
		return true
	}
	return false
}