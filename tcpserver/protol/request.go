package protol

type Message struct {
	Header Header	`json:"header"`
	Content Content `json:"content"`
}

type Header struct {
	Type string `json:"type"`
	Token string `json:"token"`
	Id string `json:"id"`
}

type Content struct {
	Method string `json:"method"`
	Params map[string]interface{} `json:"params"`
}