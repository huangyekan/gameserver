package protol

type Response struct {
	Msg string `json:"msg"`
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

var NETWORK_ERROR *Response = &Response{
	Msg:"网络出错",
	Code:900,
	Data:nil,
}


