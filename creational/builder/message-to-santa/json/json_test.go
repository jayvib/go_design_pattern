package json

import (
	"fmt"
	"go_design_pattern/creational/builder/message-to-santa"
	"testing"
)

func TestJSONMessageBuilder(t *testing.T) {
	builderDirector := &message_to_santa.BuilderDirector{
		builderMsg: new(JSONMessageBuilder),
		msg: message{
			recipient: "Santa Clause",
			text: "I will be a good boy from now on...",
		},
	}

	m, err := builderDirector.Build()
	if err != nil {
		t.Error(err)
	}

	if m.Format != "JSON" {
		t.Errorf("Expecting format to be JSON but got: %s\n", m.Format)
	}

	fmt.Println(string(m.Data))
}
