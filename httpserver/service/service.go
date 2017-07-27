package service

import (
	"net/rpc"
	"log"
	"encoding/json"
	"os"
	_"path/filepath"
	"path/filepath"
)

type ServiceConfig struct {
	Name string `json:"serviceName"`
	Url string `json:"url"`
}

var serviceMapping map[string]string = make(map[string]string)

func Init() {
	wd, _ := os.Getwd()
	file, _ := os.Open(filepath.Join(wd,"/httpserver/conf/serviceconfig.json"))
	defer file.Close()
	var serviceList []*ServiceConfig
	json.NewDecoder(file).Decode(&serviceList)
	for _, value := range serviceList {
		serviceMapping[value.Name] = value.Url
	}
}

func RemoteService(serviceName string, methodName string, params map[string]interface{}, result interface{}) error {
	client, err := rpc.DialHTTP("tcp", serviceMapping[serviceName])
	if err != nil {
		log.Println("error", err)
		return err
	}
	p, _ := json.Marshal(params)
	log.Println("rpc request [serviceName: " + serviceName + "." + methodName + " params: " + string(p) + "]")
	err = client.Call(serviceName + "." + methodName, params, result)
	if err != nil {
		log.Println(err)
		return err
	}
	r, _ := json.Marshal(result)
	log.Println("rpc response [serviceName: " + serviceName + " result: " + string(r) + "]")
	return nil
}