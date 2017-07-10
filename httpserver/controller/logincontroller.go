package controller

import "net/http"

func Login(rw http.ResponseWriter, req *http.Request)  {
	userName := req.Form["userName"]
	password := req.Form["password"]

}