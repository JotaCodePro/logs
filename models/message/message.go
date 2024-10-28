package message

type Message struct {
	Level   string `json:"level"`
	Time    string `json:"time"`
	Message string `json:"msg"`
}
