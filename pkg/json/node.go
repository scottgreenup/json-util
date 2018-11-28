package json

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

type Node struct {
	root interface{}
}

func NewFooFromRawJSON(data []byte) *Node {
	var object map[string]interface{}

	if err := json.Unmarshal(data, &object); err != nil {
		panic(err)
	}

	return &Node{
		root: object,
	}
}

func NewFooFromReader(reader io.Reader) *Node {
	input, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	return NewFooFromRawJSON(input)
}

func (f *Node) Get(key string) (*Node, error) {
	t := reflect.TypeOf(f.root)
	if t == nil {
		return nil, errors.Errorf("Unknown type of %q", key)
	}

	switch t.Kind() {
	case reflect.Map:
		jsonMap, ok := f.root.(map[string]interface{})
		if !ok {
			return nil, errors.Errorf("Unknown map, won't be able to get %q", key)
		}

		if value, ok := jsonMap[key]; ok {
			return &Node{
				root: value,
			}, nil
		}

		return nil, errors.Errorf("Map does not contain %q", key)

	case reflect.Slice:
		k, err := strconv.Atoi(key)
		if err != nil {
			return nil, errors.Errorf("Expected index for slice, not %q", key)
		}

		jsonSlice, ok := f.root.([]interface{})
		if !ok {
			return nil, errors.Errorf("Unknown slice, won't be able to get index %q", key)
		}

		if k >= len(jsonSlice) {
			return nil, errors.Errorf("Index out of bounds %q > len=%d", k, len(jsonSlice))
		}

		return &Node{
			root: jsonSlice[k],
		}, nil
	}

	return nil, errors.Errorf("Unhandled kind %q", t.Kind().String())
}

