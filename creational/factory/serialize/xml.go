package serialize

import "encoding/xml"

type XMLSerializer struct{}

func (XMLSerializer) Serialize(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}
