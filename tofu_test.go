package tofu

import (
	"encoding/json"
	"testing"
)

type Foo struct {
	Name string
}

var mapInterface = map[string]interface{}{
	"foo": map[string]interface{}{
		"foo": "one",
		"bar": "two",
		"baz": "three",
	},
	"bar": []string{
		"foo", "bar", "baz",
	},
	"baz": map[string]interface{}{
		"foo": Foo{
			Name: "name",
		},
		"bar": "two",
		"baz": []int{
			1, 2, 3,
		},
	},
}

var jsonObject = "{ \"a\": 1, \"b\": \"2\", \"c\": { \"b\": 1.2 } }"

func TestGetValue(t *testing.T) {
	path := "foo.bar"
	if Get(mapInterface, path) != "two" {
		t.Error("Get failed:", path)
	}
}

func TestGetValueFailed(t *testing.T) {
	path := "foo.bar.baz"
	if Get(mapInterface, path) != nil {
		t.Error("Get should fail:", path)
	}
}

func TestGetValueDefault(t *testing.T) {
	path := "foo.bar.baz"
	if Get(mapInterface, path, "default") != "default" {
		t.Error("Get should default:", path)
	}
}

func TestGetValueArray(t *testing.T) {
	path := "bar.0"
	if Get(mapInterface, path) != "foo" {
		t.Error("Get failed:", path)
	}
}

func TestGetValueStruct(t *testing.T) {
	path := "baz.foo.Name"
	if Get(mapInterface, path) != "name" {
		t.Error("Get failed:", path)
	}
}

func TestGetJson(t *testing.T) {
	var data map[string]interface{}
	if json.Unmarshal([]byte(jsonObject), &data) != nil {
		t.Error("Invalid JSON")
	} else {
		path := "a"
		val := GetN(data, path).Int()
		if val != 1 {
			t.Error("Get failed:", path)
		}
	}
}
