package json

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

type Handler func(path, key string, value interface{})

type Object map[string]Value

func NewObject(raw []byte) (*Object, error) {

	var jsonObject map[string]interface{}

	if err := json.Unmarshal(raw, &jsonObject); err != nil {
		return nil, err
	}

	return nil, nil
}

type Array []Value

func walk(values map[string]interface{}, handler Handler, path string, pathSuffix string) error {
	handler(path, pathSuffix, values)

	for k, v := range values {

		currentPath := fmt.Sprintf("%s.%s", path, k)

		t := reflect.TypeOf(v)
		if t == nil {
			return errors.Errorf("%q: Unknown type %+v", currentPath, v)
		}

		switch t.Kind() {
		case reflect.Map:
			jsonMap, ok := v.(map[string]interface{})
			if !ok {
				return errors.Errorf("%q: Couldn't process as map: %+v", currentPath, v)
			}

			return walk(jsonMap, handler, currentPath, k)

		case reflect.Slice:
			k, err := strconv.Atoi(k)
			if err != nil {
				return errors.Errorf("%q: Expected index for slice", currentPath)
			}

			jsonSlice, ok := v.([]interface{})
			if !ok {
				return errors.Errorf("%q: Unknown slice, won't be able to get index", currentPath)
			}

			if k >= len(jsonSlice) {
				return errors.Errorf("%q: Index out of bounds %q > len=%d", currentPath, k, len(jsonSlice))
			}

			// walk
			return nil
		}

		return errors.Errorf("Unhandled kind %q", t.Kind().String())
	}

	return nil
}
