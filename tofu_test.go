package tofu

import "testing"

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
