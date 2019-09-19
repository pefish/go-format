package go_format

import (
	"fmt"
	"testing"
)

func TestFormatClass_StructToMap(t *testing.T) {
	type Test struct {
		A uint64 `json:"a"`
		B string `json:"haha"`
	}
	test := Test{
		A: 100,
		B: `1111`,
	}
	fmt.Printf(`%#v`, Format.StructToMap(test))
	if Format.StructToMap(test)[`a`].(uint64) != 100 {
		t.Error()
	}
}

func TestFormatClass_MapToStruct(t *testing.T) {
	type Test struct {
		A uint64 `json:"a"`
		B string `json:"haha"`
	}
	test := Test{}
	Format.MapToStruct(map[string]interface{}{
		`a`:    100,
		`haha`: `1111`,
	}, &test)
	fmt.Printf(`%#v`, test)
	if test.A != 100 {
		t.Error()
	}
}

func TestFormatClass_SliceToStruct(t *testing.T) {
	type Test struct {
		A uint64 `json:"a"`
		B string `json:"haha"`
	}
	test := []Test{}
	Format.SliceToStruct([]interface{}{
		map[string]interface{}{
			`a`:    100,
			`haha`: `1111`,
		},
		map[string]interface{}{
			`a`:    100,
			`haha`: `1111`,
		},
	}, &test)
	fmt.Printf(`%#v`, test)
}
