package protol

type Response struct {
	Id string `json:"id"`
	Msg string `json:"msg"`
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

var OK *Response = &Response{
	Msg:"ok",
	Code:0,
	Data:"success",
}

var NETWORK_ERROR *Response = &Response{
	Msg:"网络出错",
	Code:900,
	Data:nil,
}


