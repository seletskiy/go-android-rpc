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
	test := server.GetViewsList("useless_layout")
	log.Printf("!!! %#v", test)

	for {
		_ = <-server.input
	}
}

func (server *rpcServer) Handle(data map[string]interface{}) {
	server.input <- data
}

func (server *rpcServer) Call(payload map[string]interface{}) map[string]interface{} {
	payloadJson, _ := json.Marshal(payload)

	resultJson := server.frontend.CallFrontend(string(payloadJson))

	var result map[string]interface{}
	err := json.Unmarshal([]byte(resultJson), &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (server *rpcServer) GetViewsList(layoutName string) map[string]interface{} {
	return server.Call(map[string]interface{}{
		"method": "ListViews",
		"layout": layoutName,
	})
}
