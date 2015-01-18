package android

import "fmt"

type Float float64

func (number Float) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%f", number)), nil
}
