package rpc

import (
	"encoding/json"
	"log"
)

var server = &rpcServer{
	input: make(chan map[string]interface{}),
}

type Frontend interface {
	CallFrontend(payload string) (reply string)
}

func Link(frontend Frontend) {
	server.frontend = frontend
	go server.Run()
}

func CallBackend(payload string) string {
	var data interface{}
	err := json.Unmarshal([]byte(payload), &data)
	if err != nil {
		log.Fatalf(`Unmarshal in Send error: %s`, err)
		return ""
	}

	server.Handle(data.(map[string]interface{}))

	// no reply for you, feeble java server!
	return ""
}
