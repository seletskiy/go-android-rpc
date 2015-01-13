package rpc

import "fmt"

type float float64

func (number float) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%f", number)), nil
}
