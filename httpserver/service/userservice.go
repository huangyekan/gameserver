package service

import (
	"github.com/huangyekan/gameserver/userservice/model"
	"log"
	"gmserver/httpserver/util/md5"
)

type UserService struct {
	
}

func (u *UserService) IsValidUser(account string, password string) bool {
	params := map[string]interface{}{"account": account}

	var user model.User
	err := RemoteService("UserService.GetUserByAccount", params, &user)
	if err != nil {
		log.Println(err)
		return false
	}

	if user.Password == md5.Encode(password) {
		return true
	}
	return false
}