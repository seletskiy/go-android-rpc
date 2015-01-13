package rpc

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/zazab/zhash"
)

type rpcServer struct {
	frontend Frontend
	input    chan map[string]interface{}
}

func (server *rpcServer) Run() {
	views := zhash.HashFromMap(server.GetViewsList("useless_layout"))

	viewId, err := views.GetString("resources", "useless_accel", "id")
	if err != nil {
		log.Printf("!!! %#v", err)
	}

	sensors := zhash.HashFromMap(server.GetSensorsList())
	sensorId, err := sensors.GetString("sensors", "TYPE_ACCELEROMETER")
	if err != nil {
		log.Printf("!!! %#v", err)
	}

	server.SubscribeToSensorValues(sensorId)

	log.Printf("!!! %#v", viewId)

	_ = server.CallViewMethod(
		viewId,
		"TextView",
		"setTextSize", 0, float(60.0),
	)

	for {
		payload := <-server.input

		data := zhash.HashFromMap(payload)

		values, _ := data.GetFloatSlice("values")

		_ = server.CallViewMethod(
			viewId,
			"TextView",
			"setText", fmt.Sprintf(
				"% 3.2f\n% 3.2f\n% 3.2f",
				values[0], values[1], values[2],
			),
		)
	}
}

func (server *rpcServer) Handle(data map[string]interface{}) {
	server.input <- data
}

func (server *rpcServer) Call(
	payload map[string]interface{},
) map[string]interface{} {
	payloadJson, _ := json.Marshal(payload)
	log.Printf("!!! SENT %#v", string(payloadJson))

	resultJson := server.frontend.CallFrontend(string(payloadJson))

	var result map[string]interface{}
	err := json.Unmarshal([]byte(resultJson), &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (server *rpcServer) GetViewsList(
	layoutName string,
) map[string]interface{} {
	return server.Call(map[string]interface{}{
		"method": "ListViews",
		"layout": layoutName,
	})
}

func (server *rpcServer) SubscribeToSensorValues(
	sensorId string,
) map[string]interface{} {
	return server.Call(map[string]interface{}{
		"method":    "SubscribeToSensorValues",
		"sensor_id": sensorId,
	})
}

func (server *rpcServer) GetSensorsList() map[string]interface{} {
	return server.Call(map[string]interface{}{
		"method": "GetSensorsList",
	})
}

func (server *rpcServer) CallViewMethod(
	id string,
	viewType string,
	methodName string,
	args ...interface{},
) map[string]interface{} {
	return server.Call(map[string]interface{}{
		"method":     "CallViewMethod",
		"id":         id,
		"type":       viewType,
		"viewMethod": methodName,
		"args":       args,
	})
}
