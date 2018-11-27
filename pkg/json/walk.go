package json

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Object struct {
	value *Value
}

func NewObject(raw []byte) (*Object, error) {
	var jsonObject map[string]interface{}

	if err := json.Unmarshal(raw, &jsonObject); err != nil {
		return nil, err
	}

	value, err := NewValue(jsonObject)
	if err != nil {
		return nil, err
	}

	return &Object{
		value: value,
	}, nil
}

func joinPath(path, suffix string) string {
	if strings.HasSuffix(path, ".") {
		return path + suffix
	}
	return path + "." + suffix
}

func joinIndex(path string, index int) string {
	if strings.HasSuffix("path", ".") {
		path = path[:len(path)-1]
	}

	return fmt.Sprintf("%s[%d]", path, index)
}

func (o *Object) Walk(handler Handler) {
	o.walk(o.value, handler, ".")
}

func (o *Object) walkObject(value *Value, handler Handler, path string) error {
	m, ok := value.Value.(map[string]interface{})
	if !ok {
		return errors.Errorf("received a non-object: %+v", value.Value)
	}

	for k, v := range m {
		currPath := joinPath(path, k)

		value, err := NewValue(v)
		if err != nil {
			return err
		}

		if err := o.walk(value, handler, currPath); err != nil {
			return err
		}
	}

	return nil
}

func (o *Object) walkArray(value *Value, handler Handler, path string) error {
	a, ok := value.Value.([]interface{})
	if !ok {
		return errors.Errorf("received a non-array: %+v", value.Value)
	}

	for i, v := range a {
		currPath := joinIndex(path, i)

		value, err := NewValue(v)
		if err != nil {
			return err
		}

		if err := o.walk(value, handler, currPath); err != nil {
			return err
		}
	}

	return nil
}

func (o *Object) walk(value *Value, handler Handler, path string) error {
	handler(path, value)

	switch value.Type {
	case ValueTypeArray:
		return o.walkArray(value, handler, path)
	case ValueTypeObject:
		return o.walkObject(value, handler, path)
	default:
		return nil
	}
}
