package rpc

import (
	"encoding/json"
	"log"
)

var server = &rpcServer{
	input: make(chan map[string]interface{}),
}

func init() {
	go server.Run()
}

type Frontend interface {
	Handle(payload string)
}

func Link(frontend Frontend) {
	server.frontend = frontend
}

func Handle(payload string) {
	var data interface{}
	err := json.Unmarshal([]byte(payload), &data)
	if err != nil {
		log.Fatalf(`Unmarshal in Send error: %s`, err)
		return
	}

	server.Handle(data.(map[string]interface{}))
}
