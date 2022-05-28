package models

import (
	"errors"
	"reflect"
	"strings"
)

func FillModel(model interface{}, data interface{}) error {
	values, success := data.(map[string]interface{})
	if !success {
		return errors.New("Error in data structure")
	}
	for key, element := range values {
		setField(model, key, element)
	}
	return nil
}

func setField(model interface{}, name string, value interface{})  {
	property := strings.Title(name)
	reflect.ValueOf(model).Elem().FieldByName(property).Set(reflect.ValueOf(value))
}

