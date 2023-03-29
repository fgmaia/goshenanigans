package main

import (
	"fmt"
)

type SourceData interface {
	map[string]any | []map[string]any
}

func Test[T SourceData](a T) *T {
	switch any(a).(type) {
	case map[string]any:
		fmt.Println("Type is map[string]interface{}")
	case []map[string]any:
		fmt.Println("Type is []map[string]interface{}")
	default:
		fmt.Println("Type is unknown")
	}
	return nil
}

func main() {
	m := map[string]any{"key": "value"}
	Test(m)
	s := []map[string]any{{"key1": "value1"}, {"key2": "value2"}}
	Test(s)
	//i := 42
	//Test(i) // compile error: i is not a MapSourceData or SliceMapSourceData
}
