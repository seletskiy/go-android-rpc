package android

import (
	"log"
	"runtime/debug"

	"github.com/zazab/zhash"
	"golang.org/x/mobile/app"
	_ "golang.org/x/mobile/bind/java"
)

// @TODO: move operations on that map to method, so generated code will not
// access it directly
var ViewTypeConstructors = map[string]func(id string) interface{}{}

type PayloadType interface{}

type SensorListener interface {
	OnChange(values []float64)

	// @TODO: not implemented
	OnAccuracyChange()
}

type OnClickListener interface {
	OnClick()
}

type OnTouchListener interface {
	OnTouch() PayloadType
}

type ViewObject interface {
	GetId() string
	GetClassName() string
}

var sensorListeners = map[string]SensorListener{}
var onClickListeners = map[string]OnClickListener{}
var onTouchListeners = map[string]OnTouchListener{}

type payloadWithReply struct {
	Data    PayloadType
	ReplyTo chan PayloadType
}

type inputChan chan payloadWithReply
type outputChan chan payloadWithReply

type backend struct {
	input   inputChan
	output  outputChan
	running bool
}

var onStart func()
var onDestroy func()

var goBackend = backend{
	input:  make(inputChan, 0),
	output: make(outputChan, 0),
}

func OnStart(callback func()) {
	onStart = callback
}

// not implemented yet
func OnDestroy(callback func()) {
	onDestroy = callback
}

func GetRequest() payloadWithReply {
	return <-goBackend.output
}

func SendEvent(payload PayloadType, replyTo chan PayloadType) {
	goBackend.input <- payloadWithReply{
		Data:    payload,
		ReplyTo: replyTo,
	}
}

func Start() {
	goBackend = backend{
		input:  make(inputChan, 0),
		output: make(outputChan, 0),
	}
	go onStart()
	goBackend.Run()
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

func CallControlMusicPlayback(
	id string,
	action string,
	loop bool,
	args ...interface{},
) map[string]interface{} {
	return goBackend.call(map[string]interface{}{
		"method":      "ControlMusicPlayback",
		"resource_id": id,
		"action":      action,
		"loop":        loop,
	})
}

func SubscribeToViewEvent(
	id string,
	viewType string,
	event string,
	callback interface{},
) map[string]interface{} {
	switch event {
	case "onClick":
		onClickListeners[id] = callback.(OnClickListener)
	case "onTouch":
		onTouchListeners[id] = callback.(OnTouchListener)
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
		view.GetId(), view.GetClassName(), "onClick",
		callback,
	)
}

func OnTouch(
	view ViewObject,
	callback OnTouchListener,
) {
	SubscribeToViewEvent(
		view.GetId(), view.GetClassName(), "onTouch",
		callback,
	)
}

func CreateView(
	id string,
	viewType string,
) ViewObject {
	goBackend.call(map[string]interface{}{
		"method": "CreateView",
		"id":     id,
		"type":   viewType,
	})

	return ViewTypeConstructors[viewType](id).(ViewObject)
}

func GetLayoutById(
	layout string,
) map[string]interface{} {
	return goBackend.call(map[string]interface{}{
		"method": "GetLayoutById",
		"layout": layout,
	})
}

func GetResourceById(
	resource string,
) string {
	response := goBackend.call(map[string]interface{}{
		"method":   "GetResourceById",
		"resource": resource,
	})
	return response["resource_id"].(string)
}

func AttachView(
	view ViewObject,
	viewGroupId string,
) {
	goBackend.call(map[string]interface{}{
		"method":      "AttachView",
		"id":          view.GetId(),
		"viewGroupId": viewGroupId,
	})
}

func ChangeLayout(
	layoutName string,
) {
	goBackend.call(map[string]interface{}{
		"method": "ChangeLayout",
		"layout": layoutName,
	})
}

func (server *backend) Run() {
	defer func() {
		err := recover()
		log.Printf("PANIC %s", err)

		log.Print(string(debug.Stack()))

		panic(err)
	}()

	server.running = true

	log.Printf("Backend started")

	for {
		event := <-server.input

		dataHash := zhash.HashFromMap(event.Data.(map[string]interface{}))
		data, err := dataHash.GetMap("data")
		if err != nil {
			panic(err)
		}

		eventName, err := dataHash.GetString("event")
		if err != nil {
			panic(err)
		}

		switch eventName {
		case "sensorChange":
			event.ReplyTo <- server.onSensorChange(zhash.HashFromMap(data))
		case "click":
			event.ReplyTo <- server.onClick(zhash.HashFromMap(data))
		case "touch":
			event.ReplyTo <- server.onTouch(zhash.HashFromMap(data))
		}
	}
}

func (server *backend) call(
	payload map[string]interface{},
) map[string]interface{} {
	replyTo := make(chan PayloadType, 0)

	server.output <- payloadWithReply{
		Data:    payload,
		ReplyTo: replyTo,
	}

	reply := <-replyTo

	return reply.(map[string]interface{})
}

func (server *backend) onSensorChange(data zhash.Hash) PayloadType {
	sensorId, err := data.GetString("sensor_id")
	if err != nil {
		panic(err)
	}

	callback, ok := sensorListeners[sensorId]
	if !ok {
		return nil
	}

	values, err := data.GetFloatSlice("values")
	if err != nil {
		panic(err)
	}

	callback.OnChange(values)

	return nil
}

func (server *backend) onClick(data zhash.Hash) PayloadType {
	viewId, err := data.GetString("viewId")
	if err != nil {
		panic(err)
	}

	callback, ok := onClickListeners[viewId]
	if !ok {
		return nil
	}

	callback.OnClick()

	return nil
}

func (server *backend) onTouch(data zhash.Hash) PayloadType {
	viewId, err := data.GetString("viewId")
	if err != nil {
		panic(err)
	}

	callback, ok := onTouchListeners[viewId]
	if !ok {
		return nil
	}

	return callback.OnTouch()
}
