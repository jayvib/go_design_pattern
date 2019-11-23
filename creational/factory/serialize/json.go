package serialize

import (
	"encoding/json"
)

type JSONSerializer struct{}

func (JSONSerializer) Serialize(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
