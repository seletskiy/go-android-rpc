package android

//go:generate ./generate.bash

import (
	"log"
	"reflect"
	"runtime/debug"

	"github.com/zazab/zhash"
	"golang.org/x/mobile/app"
	_ "golang.org/x/mobile/bind/java"
)

// @TODO: move operations on that map to method, so generated code will not
// access it directly
var ViewTypeConstructors = map[string]func(id string) interface{}{}

type PayloadType map[string]interface{}

type SensorListener interface {
	OnChange(values []float64)

	// @TODO: not implemented
	OnAccuracyChange()
}

type OnClickListener interface {
	OnClick()
}

type ViewObject interface {
	GetInternalId_() string
}

var sensorListeners = map[string]SensorListener{}
var onClickListeners = map[string]OnClickListener{}

type outputChanType struct {
	Data    PayloadType
	ReplyTo chan PayloadType
}

type inputChan chan PayloadType
type outputChan chan outputChanType

type backend struct {
	input   inputChan
	output  outputChan
	running bool
	onStart func()
}

var goBackend = backend{
	input:  make(inputChan, 0),
	output: make(outputChan, 0),
}

func OnStart(callback func()) {
	goBackend.onStart = callback
}

func GetRequest() outputChanType {
	return <-goBackend.output
}

func SendEvent(payload PayloadType) {
	goBackend.input <- payload
}

func Start() {
	if !goBackend.running {
		goBackend.Run()
	}
}

func Enter() {
	// @TODO: consider own bindings
	app.Run(app.Callbacks{})
}

func GetViewsList(
	layoutName string,
) map[string]interface{} {
	return goBackend.call(map[string]interface{}{
		"method": "ListViews",
		"layout": layoutName,
	})
}

// @TODO: create interface for Views
func GetViewById(layout, id string) interface{} {
	views := zhash.HashFromMap(GetViewsList(layout))

	viewType, err := views.GetString("resources", id, "type")
	if err != nil {
		return nil
	}

	internalId, err := views.GetString("resources", id, "id")
	if err != nil {
		return nil
	}

	// @TODO: validation via init functions in every generated code file
	return ViewTypeConstructors[viewType](internalId)
}

func SubscribeToSensorValues(
	sensorId string,
	callback SensorListener,
) map[string]interface{} {
	sensorListeners[sensorId] = callback

	return goBackend.call(map[string]interface{}{
		"method":    "SubscribeToSensorValues",
		"sensor_id": sensorId,
	})
}

func GetSensorsList() map[string]interface{} {
	return goBackend.call(map[string]interface{}{
		"method": "GetSensorsList",
	})
}

func CallViewMethod(
	id string,
	viewType string,
	methodName string,
	args ...interface{},
) map[string]interface{} {
	return goBackend.call(map[string]interface{}{
		"method":     "CallViewMethod",
		"id":         id,
		"type":       viewType,
		"viewMethod": methodName,
		"args":       args,
	})
}

func SubscribeToViewEvent(
	id string,
	viewType string,
	event string,
	callback OnClickListener,
) map[string]interface{} {
	switch event {
	case "onClick":
		onClickListeners[id] = callback
	}

	return goBackend.call(map[string]interface{}{
		"method": "SubscribeToViewEvent",
		"id":     id,
		"type":   viewType,
		"event":  event,
	})
}

func OnClick(
	view ViewObject,
	callback OnClickListener,
) {
	SubscribeToViewEvent(
		view.GetInternalId_(), reflect.TypeOf(view).Name(), "onClick",
		callback,
	)
}

func (server *backend) Run() {
	defer func() {
		err := recover()
		log.Printf("PANIC %s", err)

		log.Print(string(debug.Stack()))

		panic(err)
	}()

	server.running = true

	if server.onStart != nil {
		server.onStart()
	}

	log.Printf("Backend started")

	for {
		payload := <-server.input

		event := zhash.HashFromMap(payload)
		data, err := event.GetMap("data")
		if err != nil {
			panic(err)
		}

		eventName, err := event.GetString("event")
		if err != nil {
			panic(err)
		}

		switch eventName {
		case "sensorChange":
			server.onSensorChange(zhash.HashFromMap(data))
		case "click":
			server.onClick(zhash.HashFromMap(data))
		}
	}
}

func (server *backend) call(
	payload map[string]interface{},
) map[string]interface{} {
	replyTo := make(chan PayloadType, 0)

	server.output <- outputChanType{
		Data:    payload,
		ReplyTo: replyTo,
	}

	return <-replyTo
}

func (server *backend) onSensorChange(data zhash.Hash) {
	sensorId, err := data.GetString("sensor_id")
	if err != nil {
		panic(err)
	}

	callback, ok := sensorListeners[sensorId]
	if !ok {
		return
	}

	values, err := data.GetFloatSlice("values")
	if err != nil {
		panic(err)
	}

	callback.OnChange(values)
}

func (server *backend) onClick(data zhash.Hash) {
	viewId, err := data.GetString("view_id")
	if err != nil {
		panic(err)
	}

	callback, ok := onClickListeners[viewId]
	if !ok {
		return
	}

	callback.OnClick()
}
