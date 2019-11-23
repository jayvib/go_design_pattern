package serialize

import "github.com/pkg/errors"

type format int

const (
	JSON format = 0
	XML  format = 1
)

type Serializer interface {
	// Serialize will marshal the 'v' into an array of bytes
	// thats been serialize
	Serialize(v interface{}) ([]byte, error)
}

func NewSerializer(f format) (Serializer, error) {
	switch f {
	case JSON:
		return new(JSONSerializer), nil
	case XML:
		return new(XMLSerializer), nil
	default:
		return nil, errors.New("Format doesn't recognized.")
	}
}
