package rpc

import (
	"encoding/json"
	"log"
)

type rpcServer struct {
	frontend Frontend
	input    chan map[string]interface{}
}

func (server *rpcServer) Run() {
	for {
		data := <-server.input
		log.Printf("!!! %#v", data)
	}
}

func (server *rpcServer) Handle(data map[string]interface{}) {
	server.input <- data
}

func (server *rpcServer) hideView(id string) {
	payload, _ := json.Marshal(map[string]interface{}{
		"id":     id,
		"action": "hide",
	})

	server.frontend.Handle(string(payload))
}

func (server *rpcServer) showView(id string) {
	payload, _ := json.Marshal(map[string]interface{}{
		"id":     id,
		"action": "show",
	})

	server.frontend.Handle(string(payload))
}
