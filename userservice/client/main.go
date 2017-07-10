package main

import (
	"net/rpc"
	"log"
	"gmserver/userservice/model"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9001")
	if err != nil {
		log.Println("error", err)
	}
	params := map[string]interface{}{"account":"111"}
	result := model.User{}
	err = client.Call("UserService.GetUserByAccount", params, &result)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
}