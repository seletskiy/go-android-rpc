package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime/debug"

	"github.com/seletskiy/go-android-rpc/android"
	"github.com/zazab/zhash"

	"github.com/seletskiy/go-android-rpc/android/sdk"

	// required for linking
	_ "github.com/seletskiy/go-android-rpc/android/modules"
)

type AccelerometerHandler struct {
	text sdk.TextView
}

func (handler AccelerometerHandler) OnChange(values []float64) {
	handler.text.SetText1s(fmt.Sprintf(
		"% 3.2f\n% 3.2f\n% 3.2f", values[0], values[1], values[2],
	))
}

func (handler AccelerometerHandler) OnAccuracyChange() {}

type ButtonHandler struct {
	button          sdk.Button
	accelerometerId string
}

func (handler ButtonHandler) OnClick() {
	texts := []string{
		"oh, hellu there!",
		"you are my frrrend",
		"wow, it is you agan",
		"wow, really wow!",
	}

	handler.button.PerformHapticFeedback(0)
	handler.button.SetText1s(texts[rand.Intn(len(texts))])

	result, err := handler.button.IsShown()
	log.Printf("%#v", result)
	log.Printf("%#v", err)

	android.OpenWebPage("http://github.com/seletskiy/go-android-rpc/")

	android.UnsubscribeToSensorValues(
		handler.accelerometerId,
	)

	//android.ChangeLayout("another_layout")
}

func (handler ButtonHandler) OnTouch() android.PayloadType {
	texts := []string{
		`/\/\/\/\/\/\/\/\`,
		`\/\/\/\/\/\/\/\/`,
	}

	handler.button.PerformHapticFeedback(0)

	handler.button.SetText1s(texts[rand.Intn(len(texts))])

	// wow, it will be returned as OnTouch result in Java!
	return false
}

func start() {
	sensors := zhash.HashFromMap(android.GetSensorsList())

	accelDisplay := android.GetViewById("useless_accel").(sdk.TextView)
	accelDisplay.SetTextSize(40.0)

	layout := zhash.HashFromMap(android.GetLayoutById("main_layout"))
	layout_id, err := layout.GetString(
		"layout_id",
	)
	if err != nil {
		panic(err)
	}

	accelerometerId, err := sensors.GetString(
		"sensors", "TYPE_ACCELEROMETER",
	)

	touchButton := android.GetViewById("useless_touch_button").(sdk.Button)
	android.OnTouch(touchButton, ButtonHandler{touchButton, accelerometerId})

	if err != nil {
		panic(err)
	}

	android.SubscribeToSensorValues(
		accelerometerId,
		AccelerometerHandler{accelDisplay},
	)

	newView := android.CreateView("123", "android.widget.Button").(sdk.Button)
	newView.SetText1s("I'm generated!")

	android.OnClick(newView, ButtonHandler{newView, accelerometerId})
	android.AttachView(newView, layout_id)
	a := android.GetViewById("123")
	log.Printf("main.go:102 %#v", a)

	newTextEdit := android.CreateView("124", "android.widget.EditText").(sdk.EditText)
	android.AttachView(newTextEdit, layout_id)
	newTextEdit.SetText1s("123123")

}

func main() {
	defer func() {
		err := recover()
		log.Printf("PANIC %s", err)

		log.Print(string(debug.Stack()))
	}()

	android.OnStart(start)
	android.Enter()
}
