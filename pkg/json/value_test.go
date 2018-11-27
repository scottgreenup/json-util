package json

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNull(t *testing.T) {
	data := []byte(`{"a": null}`)
	var object map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &object))

	v, err := NewValue(object["a"])
	require.NoError(t, err)

	assert.Equal(t, v.Type, ValueTypeNull)
	assert.Nil(t, v.Value)
}

func TestInteger(t *testing.T) {
	data := []byte(`{"a": 42}`)
	var object map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &object))

	v, err := NewValue(object["a"])
	require.NoError(t, err)

	assert.Equal(t, v.Type, ValueTypeNumber)
	assert.Equal(t, v.Value.(float64), 42.0)
}

func TestString(t *testing.T) {
	data := []byte(`{"a": "runyoufools"}`)
	var object map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &object))

	v, err := NewValue(object["a"])
	require.NoError(t, err)

	assert.Equal(t, v.Type, ValueTypeString)
	assert.Equal(t, v.Value.(string), "runyoufools")
}

func TestBoolean(t *testing.T) {
	data := []byte(`{"a": true}`)
	var object map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &object))

	v, err := NewValue(object["a"])
	require.NoError(t, err)

	assert.Equal(t, v.Type, ValueTypeBoolean)
	assert.Equal(t, v.Value.(bool), true)
}

func TestArray(t *testing.T) {
	data := []byte(`{"a": [1,2,3,4]}`)
	var object map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &object))

	v, err := NewValue(object["a"])
	require.NoError(t, err)

	assert.Equal(t, v.Type, ValueTypeArray)
	assert.Equal(t, v.Value.([]interface{}), []interface{}{1.0, 2.0, 3.0, 4.0})
}

func TestObject(t *testing.T) {
	data := []byte(`{"a": {"b": 4}}`)
	var object map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &object))

	v, err := NewValue(object["a"])
	require.NoError(t, err)

	assert.Equal(t, v.Type, ValueTypeObject)
	assert.Equal(t, v.Value.(interface{}), map[string]interface{}{"b": 4.0})
}
