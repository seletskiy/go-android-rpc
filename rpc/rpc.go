package rpc

import (
	"encoding/json"
	"fmt"
	"time"
)

type rpcServer struct {
	frontend Frontend
	input    chan map[string]interface{}
}

func (server *rpcServer) Run() {
	for {
		data := <-server.input
		switch data["event"] {
		case "onClick":
			// from R.java
			// @TODO: create goguibind
			if data["id"] == fmt.Sprintf("%d", 0x7f030000) {
				server.hideView(data["id"].(string))
				time.Sleep(time.Second * 1)
				server.showView(data["id"].(string))
			}
		}
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
