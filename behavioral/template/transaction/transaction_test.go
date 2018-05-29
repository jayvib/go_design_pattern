package transaction

import (
	"strings"
	"testing"
)

func TestTemplate_ExecuteAlgorithm(t *testing.T) {
	t.Run("Using interface", func(t *testing.T) {
		s := &TestStruct{}
		s.Template = &TemplateImpl{}
		res := s.ExecuteAlgorithm(s)
		expected := "world"

		if !strings.Contains(res, expected) {
			t.Errorf("Expected string '%s' wasn't found on returned string\n", expected)
		}
	})

	t.Run("Using First Class function adapter to an interface", func(t *testing.T) {
		s := &TestStruct{}
		msgRetreiver := NewMessageRetrieverAdapter(func() string { // using adapter pattern
			return "world"
		})

		if msgRetreiver == nil {
			t.Fatal("Message Retriever must not be nill it will panic ahead")
		}

		s.Template = &AnonymousTemplate{}
		res := s.ExecuteAlgorithm(msgRetreiver)
		expected := "world"

		if !strings.Contains(res, expected) {
			t.Errorf("Expected string '%s' wasn't found on returned string\n", expected)
		}
	})
}
