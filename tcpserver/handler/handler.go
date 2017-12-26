package handler

import (
	"strings"
	"gmserver/tcpserver/protol"
	"log"
	"net/rpc"
)


func Dispatcher(message *protol.Message) *protol.Response {
	interfaceName, methodName := parseContent(&message.Content)
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9002")
	if err != nil {
		log.Println("error", err)
		return &protol.Response{
			Msg:  err.Error(),
			Code: 1000,
		}
	}
	var reply interface{}
	err = client.Call(interfaceName+"."+methodName, message.Content.Params, &reply)
	if err != nil {
		return &protol.Response{
			Msg:  err.Error(),
			Code: 1000,
		}
	}
	return buildResopnse(reply, message.Header.Id)
}

func parseContent(content *protol.Content) (interfaceName string, methodName string) {
	values := strings.Split(content.Method, ".")
	return values[0], values[1]
}

func buildResopnse(reply interface{}, id string) *protol.Response {
	return &protol.Response{
		Id: id,
		Code: 0,
		Data: reply,
	}
}
