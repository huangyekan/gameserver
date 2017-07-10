package service

import (
	"net/rpc"
	"log"
)

func RemoteService(serviceName string, params map[string]interface{}, result interface{}) error {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9001")
	if err != nil {
		log.Println("error", err)
		return err
	}
	err = client.Call(serviceName, params, result)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}