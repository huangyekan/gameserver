package controller

import "net/http"

func Login(rw http.ResponseWriter, req *http.Request)  {
	account := req.Form["account"]
	password := req.Form["password"]

}