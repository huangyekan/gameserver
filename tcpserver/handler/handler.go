package handler

import (
	"strings"
	"gmserver/tcpserver/protol"
	"log"
	"net/rpc"
)

func Dispatcher(content *protol.Content) *protol.Response {
	interfaceName, methodName := parseContent(content)

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9001")
	if err != nil {
		log.Println("error", err)
		return &protol.Response{
			Msg:  err.Error(),
			Code: 1000,
		}
	}
	var reply interface{}
	err = client.Call(interfaceName+"."+methodName, content.Params, &reply)
	if err != nil {
		return &protol.Response{
			Msg:  err.Error(),
			Code: 1000,
		}
	}
	return buildResopnse(reply)
}

func parseContent(content *protol.Content) (interfaceName string, methodName string) {
	values := strings.Split(content.Method, ".")
	return values[0], values[1]
}

func buildResopnse(reply interface{}) *protol.Response {
	return &protol.Response{
		Code: 100,
		Data: reply,
	}
}
