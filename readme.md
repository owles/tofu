# Tofu

Simple golang collection helper for get value in an `interface{}` by `value.path`.

### Install and usage

``` 
go get github.com/owles/tofu
```

For get value use `.` separated pah.

```go
myMap := map[string]interface{}{
    "foo": map[string]interface{}{
        "bar": []string{"apple", "orange", "potato"},	
    }
}

val := tofu.Get(myMap, "foo.bar.0", "default")
fmt.Println(val) // apple

val = tofu.Get(myMap, "foo.bar.baz")
fmt.Println(val) // nil
```