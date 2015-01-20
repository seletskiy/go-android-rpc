package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/seletskiy/go-android-rpc/android"

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
	log.Printf("!!! go here")

	result := handler.button.GetBaseline()
	log.Printf("!!! go there")
	log.Printf("!!! go result %v", result["result"])
}

func (handler ButtonHandler) OnTouch() {
	texts := []string{
		"/\\/\\/\\/\\/\\/\\/\\/\\",
		"\\/\\/\\/\\/\\/\\/\\/\\/",
	}

	handler.button.PerformHapticFeedback(0)

	handler.button.SetText1s(texts[rand.Intn(len(texts))])
}

func start() {
	uselessButton := android.GetViewById("useless_layout", "useless_button").(sdk.Button)
	android.OnClick(uselessButton, ButtonHandler{uselessButton})
}

func main() {
	android.OnStart(start)
	android.Enter()
}
