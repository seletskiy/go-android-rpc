package main

import (
	"fmt"
	"log"
	"math/rand"

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
	button sdk.Button
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

	result := handler.button.GetText()
	log.Printf("!!! current text: %s", result)

	android.ChangeLayout("another_layout")
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
	log.Printf("%#v", "onStart")
	sensors := zhash.HashFromMap(android.GetSensorsList())

	accelDisplay := android.GetViewById("main_layout", "useless_accel").(sdk.TextView)
	accelDisplay.SetTextSize(40.0)

	layout := zhash.HashFromMap(android.GetLayoutById("main_layout"))
	layout_id, err := layout.GetString(
		"layout_id",
	)
	if err != nil {
		panic(err)
	}

	touchButton := android.GetViewById("main_layout", "useless_touch_button").(sdk.Button)
	android.OnTouch(touchButton, ButtonHandler{touchButton})

	accelerometerId, err := sensors.GetString(
		"sensors", "TYPE_ACCELEROMETER",
	)

	if err != nil {
		panic(err)
	}

	android.SubscribeToSensorValues(
		accelerometerId,
		AccelerometerHandler{accelDisplay},
	)

	newView := android.CreateView("123", "android.widget.Button").(sdk.Button)
	log.Printf("%#v", newView)
	newView.SetText1s("I'm generated!")

	android.OnClick(newView, ButtonHandler{newView})
	android.AttachView(newView, layout_id)

	newTextEdit := android.CreateView("124", "android.widget.EditText").(sdk.EditText)
	android.AttachView(newTextEdit, layout_id)
	newTextEdit.SetText1s("123123")

	log.Printf("%#v", "onStart END")
}

func main() {
	android.OnStart(start)
	android.Enter()
}
