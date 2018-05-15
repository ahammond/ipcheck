package ipcheck

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNewTextAccumulator(t *testing.T) {
	expected := "a\nb\nc\n"
	a := NewTextAccumulator()
	a.Append("junk", "a")
	a.Append("junk", "b")
	a.Append("junk", "c")
	r := a.Read().String()
	if r != expected {
		t.Fatalf(
			"\nexpected: %#v\ngot:      %#v",
			expected, r)
	}
}

func TestNewJsonAccumulator(t *testing.T) {
	expected := map[string]string{
		"k": "v",
		"a": "b",
	}
	a := NewJsonAccumulator()
	for k, v := range expected {
		a.Append(k, v)
	}
	r := make(map[string]string)
	json.Unmarshal(a.Read().Bytes(), &r)
	if !reflect.DeepEqual(r, expected) {
		t.Fatalf(
			"\nexpected: %#v\ngot:      %#v",
			expected, r)
	}
}
