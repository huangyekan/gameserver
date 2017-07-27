package controller

import (
	"net/http"
	"gmserver/httpserver/service"
	"log"
	"text/template"
	"os"
	"path/filepath"
)

var userService *service.UserService = new(service.UserService)




func Index(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	cookie, err := r.Cookie("token")
	wd, _ := os.Getwd()
	templatePath := filepath.Join(wd, "/httpserver/static/template/")
	if err != nil {
		log.Println("get cookie error", err)
		log.Println(filepath.Join(templatePath, "index.html"))
		t, _ := template.ParseFiles(filepath.Join(templatePath, "index.html"))
		t.Execute(w, nil)
		return
	}
	token := cookie.Value
	if !userService.CheckToken(token) {
		t, _ := template.ParseFiles(filepath.Join(templatePath, "index.html"))
		t.Execute(w, nil)
		return
	}
}


func Login(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	account := r.Form["account"]
	password := r.Form["password"]
	log.Println(account)
	log.Println(password)
 	if !userService.IsValidUser(account[0], password[0]) {
		w.Write([]byte("can not login"))
	}
	w.Write([]byte("is login"))
}


func RegisterIndex(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()

}

func Register(w http.ResponseWriter, r *http.Request)  {
	
}