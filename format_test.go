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
		A: 568262448,
		B: `1111`,
	}
	fmt.Println(fmt.Sprintf(`%+v`, Format.StructToMap(test)))
	if Format.StructToMap(test)[`a`].(float64) != 568262448 {
		t.Error()
	}
}
