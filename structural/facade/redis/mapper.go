package redis

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func getValue(value reflect.Value) reflect.Value {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}

func ObjectMapper(obj interface{}) map[string]interface{} {
	value := getValue(reflect.ValueOf(obj))
	typeOf := value.Type()
	out := make(map[string]interface{})
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		// Get the tag
		redisTag := field.Tag.Get("redis")

		var fieldValue string
		switch value.Field(i).Kind() {
		case reflect.String:
			fieldValue = value.Field(i).Interface().(string)
		case reflect.Int8:
			int8Value := value.Field(i).Interface().(int8)
			fieldValue = fmt.Sprintf("%d", int8Value)
		case reflect.Float64:
			float64Value := value.Field(i).Interface().(float64)
			fieldValue = strconv.FormatFloat(float64Value, 'f', 2, 64)
		}

		out[redisTag] = fieldValue
	}

	return out
}

func MapParser(m map[string]string, storeTo interface{}) error {
	if k := reflect.TypeOf(storeTo).Kind(); k != reflect.Ptr {
		return errors.New("store to not a pointer")
	}

	value := reflect.ValueOf(storeTo).Elem()

	if value.Kind() != reflect.Struct {
		return errors.New("underlying type of store to is not a struct")
	}

	typeOf := value.Type()

	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)

		redisTag, ok := field.Tag.Lookup("redis")
		if !ok {
			continue
		}

		// check if the tag is in m
		v, ok := m[redisTag]
		if !ok {
			continue
		}

		// what is the underlying type of the field?
		switch field.Type.Kind() {
		case reflect.String:
			value.Field(i).SetString(v)
		case reflect.Int8:
			int8Val, _ := strconv.ParseInt(v, 10, 8)
			value.Field(i).SetInt(int8Val)
		case reflect.Float64:
			float64Val, _ := strconv.ParseFloat(v, 64)
			value.Field(i).SetFloat(float64Val)
		default:
			continue
		}
	}
}