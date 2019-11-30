package worker

import (
	"testing"
)


func dummHandler(d interface{}) (interface{}, error) {
	return d, nil
}

func TestDispatcher(t *testing.T) {
	d := NewDispatcher(3)

	d.Dispatch(dummHandler)

	inputs := map[string]bool{"luffy":true, "zoro":true, "sanji":true, "nami": true, "chopper": true}
	for in := range inputs {
		d.MakeRequest(Request{in})
	}

	d.MakeRequest(Request{"robin"})

	d.Stop()
}
