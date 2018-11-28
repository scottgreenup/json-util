package json

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPaths(t *testing.T) {
	data := []byte(`{
		"a": {
			"b": { 
				"c": [{
					"A": 100,
					"B": 200
				}, 15]
			},
			"d": { 
				"e": 1
			}
		}
}`)

	handler := func(path string, value *Value) {
		fmt.Printf("%s: %+v\n", path, value)
	}

	object, err := NewObject(data)
	require.NoError(t, err)

	require.NoError(t, object.Walk(handler))
}

func StringSliceDiff(x []string, y []string) []string {
	xminusy := make([]string, 0, len(x))
	for i, _ := range x {
		found := false
		for j, _ := range y {
			if x[i] == y[j] {
				found = true
				break
			}
		}

		if !found {
			xminusy = append(xminusy, x[i])
		}
	}

	return xminusy
}

func TestPathEnumeration(t *testing.T) {

	tester := func(rawJSON []byte, expectedPaths []string) {
		foundPaths := []string{}

		handler := func(path string, value *Value) {
			foundPaths = append(foundPaths, path)
		}

		object, err := NewObject(rawJSON)
		require.NoError(t, err)
		require.NoError(t, object.Walk(handler))

		extraFound := StringSliceDiff(foundPaths, expectedPaths)
		require.Len(t, extraFound, 0, "%+v", extraFound)

		notFound := StringSliceDiff(expectedPaths, foundPaths)
		require.Len(t, notFound, 0, "+v", notFound)
	}

	t.Run("basic-depth", func(t *testing.T) {
		data := []byte(`{"a": {"b": {"c": "foo"}}}`)
		expectedPaths := []string{
			".",
			".a",
			".a.b",
			".a.b.c",
		}
		tester(data, expectedPaths)
	})

	t.Run("basic-breadth", func(t *testing.T) {
		data := []byte(`
			{
				"a": {
					"b": 100,
					"c": 200
				},
				"d": {
					"e": 1000,
					"f": 2000
				},
				"g": {
					"h": 10000,
					"i": 20000
				}
			}
		`)
		expectedPaths := []string{
			".",
			".a",
			".a.b",
			".a.c",
			".d",
			".d.e",
			".d.f",
			".g",
			".g.h",
			".g.i",
		}

		tester(data, expectedPaths)
	})
}
