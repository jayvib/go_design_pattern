package serialize

import (
	"encoding/xml"
	"github.com/kubernetes/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"
	"testing"
)

type Person struct {
	FirstName string
	LastName  string
}

func TestNewSerializer(t *testing.T) {
	data := []Person{
		{"Jayson", "Vibandor"},
		{"Jimboy", "Vibandor"},
		{"Jumily", "Vibandor"},
	}

	t.Run("JSONSerializer", func(t *testing.T) {
		jsonSerializer, err := NewSerializer(JSON)
		if err != nil {
			t.Fatal(err)
		}

		for _, d := range data {
			b, err := jsonSerializer.Serialize(&d)
			if err != nil {
				t.Fatal(err)
			}
			if len(b) == 0 {
				t.Error("Expecting the marshal bytes not empty")
			}

			out := struct {
				FirstName string
				LastName  string
			}{}

			err = json.Unmarshal(b, &out)
			if err != nil {
				t.Fatal(err)
			}

			if d.FirstName != out.FirstName {
				t.Error("Expected firstname must be matched but got:", out.FirstName)
			}

			if d.LastName != out.LastName {
				t.Error("Expected lastname must be matched but got:", out.LastName)
			}
		}
	})

	t.Run("XMLSerializer", func(t *testing.T) {
		xmlSerializer, err := NewSerializer(XML)
		if err != nil {
			t.Fatal(err)
		}

		for _, d := range data {
			b, err := xmlSerializer.Serialize(&d)
			if err != nil {
				t.Fatal(err)
			}

			if len(b) == 0 {
				t.Error("expecting marshaled xml byte must non-zero")
			}

			out := Person{}

			err = xml.Unmarshal(b, &out)
			if err != nil {
				t.Error(err)
			}
		}
	})

}
