package model

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}
