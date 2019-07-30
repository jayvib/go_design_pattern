package singleton

import (
	"testing"
	"reflect"
)


func TestGetInstance(t *testing.T) {
	s := GetInstance()
	s.AddOne()	

	concrete1 := s.(*Count)
	if concrete1.count != 1 {
		t.Error("After the first add the count value should be 1")
	}

	s2 := GetInstance()
	s2.AddOne()

	concrete2 := s.(*Count)
	if concrete2.count != 2 {
		t.Error("After the second add the count value should be 2")
	}

	if !reflect.DeepEqual(s, s2) {
		t.Error("Expecting first instance is the same with the second instance")
	}
	
}
