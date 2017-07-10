package controller

import (
	"net/http"
	"gmserver/httpserver/service"
	"log"
)

func Login(rw http.ResponseWriter, req *http.Request)  {
	req.ParseForm()
	account := req.Form["account"]
	password := req.Form["password"]
	userService := new(service.UserService)
	log.Println(account)
	log.Println(password)
	isUserValid :=  userService.IsValidUser(account[0], password[0])
	if (!isUserValid) {
		rw.Write([]byte("user not exist"))
	}
	rw.Write([]byte("is login"))
}