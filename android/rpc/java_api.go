package rpc

import (
	"encoding/json"
	"log"
	"runtime/debug"

	"github.com/seletskiy/go-android-rpc/android"
)

type Frontend interface {
	CallFrontend(payload string) (reply string)
}

func StartBackend(javaFrontend Frontend) {
	go android.Start()

	go func() {
		defer func() {
			err := recover()
			log.Printf("PANIC %s", err)

			log.Print(string(debug.Stack()))

			panic(err)
		}()

		for {
			request := android.GetRequest()

			jsonRequest, err := json.Marshal(request.Data)
			if err != nil {
				panic(err)
			}

			jsonResult := javaFrontend.CallFrontend(string(jsonRequest))

			var result android.PayloadType
			err = json.Unmarshal([]byte(jsonResult), &result)
			if err != nil {
				log.Printf(
					"Error while parsing reply: %s; reply: '%s'",
					err, jsonResult,
				)
			}

			request.ReplyTo <- result
		}
	}()
}

func CallBackend(payload string) string {
	var data android.PayloadType
	err := json.Unmarshal([]byte(payload), &data)
	if err != nil {
		log.Fatalf(`Unmarshal in Send error: %s`, err)
		return ""
	}

	android.SendEvent(data)

	// no reply for you, feeble java!
	return ""
}
