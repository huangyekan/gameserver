package server

import (
	"net/http"
	"log"
	"gmserver/httpserver/controller"
	"gmserver/httpserver/service"
)

func init()  {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	service.Init()
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/index", controller.Index)
}

func Run(port string) {
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
