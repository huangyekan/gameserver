package server

import (
	"net/http"
	"log"
	"gmserver/httpserver/controller"
)

func init()  {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	http.HandleFunc("/login", controller.Login)
}

func Run(port string) {
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
