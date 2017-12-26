package server

import (
	"net"
	"log"
	"encoding/json"
	"gmserver/tcpserver/protol"
	"gmserver/tcpserver/handler"
	"gmserver/tcpserver/util"
)

var ConnMap = make(map[string]*net.TCPConn)

const (
	PING     = "ping"
	REQUEST  = "request"
	REGISTER = "register"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func Run(port int) {
	lisener, err := net.ListenTCP("tcp",
		&net.TCPAddr{
			IP:   net.ParseIP("0.0.0.0"),
			Port: port,
		})
	if err != nil {
		log.Fatal("服务启动异常", err.Error())
	}
	for {
		conn, err := lisener.AcceptTCP();
		log.Println("accept")
		if err != nil {
			log.Println("连接失败", err.Error())
			continue
		}
		go handleConnection(conn)
		defer conn.Close()
	}
}

func handleConnection(conn *net.TCPConn) {
	buf := make([]byte, 1024)
	tmpBuf := make([]byte, 0)
	var messageChannel = make(chan string, 128)
	go handleMessage(conn, messageChannel)
	for {
		n, err := conn.Read(buf)
		//buf := append(buf[:n], tmpBuf...)
		buf := append(tmpBuf, buf[:n]...)
		if err != nil {
			log.Println("读取数据异常", err.Error())
			break;
		}
		messages, leftBuf := protol.Decode(buf, []string{})
		for _, message := range messages {
			log.Println("message : ", message)
			messageChannel <- message
		}
		tmpBuf = leftBuf
	}

}

func handleMessage(conn *net.TCPConn, messageChannel chan string) {
	for {
		message := <-messageChannel
		msg := parseMessage(message)
		switch msg.Header.Type {
		case PING:
			ping(conn, msg)
		case REQUEST:
			doRequest(msg)
		case REGISTER:
			doRegister(conn, msg)
		}
	}
}

func responseMessage(response *protol.Response, conn *net.TCPConn) {
	res, _ := json.Marshal(response)
	log.Println("response ", string(res))
	res = append(util.IntToBytes(len(string(res))), res...)
	conn.Write(res)
}

func doRequest(message *protol.Message) {
	response := handler.Dispatcher(message)
	conn := ConnMap[message.Header.Token]
	if conn == nil {
		log.Println("用户已断开")
		return
	}
	log.Println(conn.RemoteAddr().String())
	responseMessage(response, conn)
}

func doRegister(conn *net.TCPConn, message *protol.Message) {
	log.Println(" register ", "ip : ", conn.RemoteAddr().String(), "token : ", message.Header.Token)
	ConnMap[message.Header.Token] = conn
	response := protol.OK
	response.Id = message.Header.Id
	responseMessage(response, conn)
}

func ping(conn *net.TCPConn, message *protol.Message) {
	log.Println(conn.RemoteAddr().String() + " ping success")
	response := protol.OK
	response.Id = message.Header.Id
	responseMessage(response, conn)
}

func parseMessage(message string) *protol.Message {
	result := new(protol.Message)
	err := json.Unmarshal([]byte(message), &result)
	if err != nil {
		log.Println("json解析异常", err.Error())
	}
	return result
}
