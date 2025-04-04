package protocol

import (
	"fmt"
	"reflect"
)

func CallHandler(handler interface{}, params ...interface{}) ([]interface{}, error) {
	handlerValue := reflect.ValueOf(handler)
	handlerType := handlerValue.Type()

	if len(params) != handlerType.NumIn() {
		return nil, fmt.Errorf("handler expects %d parameters, got %d", handlerType.NumIn(), len(params))
	}

	args := make([]reflect.Value, len(params))
	for i, p := range params {
		args[i] = reflect.ValueOf(p)
	}

	results := handlerValue.Call(args)

	var err error
	if len(results) > 0 {
		lastResult := results[len(results)-1]
		errorInterface := reflect.TypeOf((*error)(nil)).Elem()
		if lastResult.Type().Implements(errorInterface) {
			// If the last value is non-nil, assign it to err.
			if !lastResult.IsNil() {
				err = lastResult.Interface().(error)
			}
			// Remove the error value from results.
			results = results[:len(results)-1]
		}
	}

	out := make([]interface{}, len(results))
	for i, r := range results {
		out[i] = r.Interface()
	}

	return out, err
}
