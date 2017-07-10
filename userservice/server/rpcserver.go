package server

import (
	"net/rpc"
	"gmserver/userservice/service"
	"net"
	"net/http"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	rpc.Register(new(service.Echo))
	rpc.Register(new(service.UserService))
}

func Run(port string) {
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("error", err)
	}
	http.Serve(l, nil)
}
