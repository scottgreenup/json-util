package json

import (
	"reflect"

	"github.com/pkg/errors"
)

type Value struct {
	Type  ValueType
	Value interface{}
}

type ValueType uint64

const (
	ValueTypeString ValueType = iota
	ValueTypeNumber
	ValueTypeObject
	ValueTypeArray
	ValueTypeBoolean
	ValueTypeNull
)

func NewValue(v interface{}) (*Value, error) {
	if v == nil {
		return &Value{
			Type:  ValueTypeNull,
			Value: nil,
		}, nil
	}

	// TODO: It might be possible that this will only return nil when v is nil. In that case, the previous
	// conditional statement checking whether v is nil can be collated into the lookup-table below.
	t := reflect.TypeOf(v)
	if t == nil {
		return nil, errors.Errorf("Unknown type of value: %+v", v)
	}

	// TODO: Are all numbers just float64?
	lut := map[reflect.Kind]ValueType{
		reflect.Bool:    ValueTypeBoolean,
		reflect.Float32: ValueTypeNumber,
		reflect.Float64: ValueTypeNumber,
		reflect.Int:     ValueTypeNumber,
		reflect.Int8:    ValueTypeNumber,
		reflect.Int16:   ValueTypeNumber,
		reflect.Int32:   ValueTypeNumber,
		reflect.Int64:   ValueTypeNumber,
		reflect.Map:     ValueTypeObject,
		reflect.Slice:   ValueTypeArray,
		reflect.String:  ValueTypeString,
		reflect.Uint:    ValueTypeNumber,
		reflect.Uint8:   ValueTypeNumber,
		reflect.Uint16:  ValueTypeNumber,
		reflect.Uint32:  ValueTypeNumber,
		reflect.Uint64:  ValueTypeNumber,
	}

	vt, ok := lut[t.Kind()]
	if !ok {
		return nil, errors.Errorf("Unhandled type of value: %+v (%s)", v, t.Kind().String())
	}

	return &Value{
		Type:  vt,
		Value: v,
	}, nil
}
