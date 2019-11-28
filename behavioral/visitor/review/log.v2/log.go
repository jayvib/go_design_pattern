package log

import (
	"io"
	"reflect"
)

type WriteLogger struct {
	Msg string
	w   io.Writer
}

func (cl *WriteLogger) Log() {
	cl.w.Write([]byte(cl.Msg))
}

func (cl *WriteLogger) SetMsg(msg string) {
	cl.Msg = msg
}

func (cl *WriteLogger) Accept(v Visitor) {
	v.Visit(cl)
}

type MsgSetter interface {
	SetMsg(string)
}

// This acts as an Visitor
type Visitor interface {
	Visit(interface{})
}

type visitorImpl struct {
	Msg string
}

func (v *visitorImpl) Visit(obj interface{}) {
	//if ms, ok := obj.(MsgSetter); ok {
	//	ms.SetMsg(v.Msg)
	//}

	// Process only the Pointer Type
	val := reflect.ValueOf(obj)
	switch val.Kind() {
	case reflect.Ptr:
		// Modify the object
		concVal := val.Elem()
		switch concVal.Kind() { // check the type if struct
		case reflect.Struct:
			// Modify the field
			concType := concVal.Type()
			for i := 0; i < concType.NumField(); i++ {
				f := concType.Field(i)
				if f.Name == "Msg" {
					concVal.Field(i).SetString(v.Msg)
				}
			}
		}
	}
}
