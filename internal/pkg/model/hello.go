package model

type HelloResponse struct {
	Message string `json:"message"`
}

type AllHelloResponse struct {
	Messages []string `json:"messages"`
}
