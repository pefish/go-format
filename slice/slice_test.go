package go_format_slice

import (
	"fmt"
	"testing"

	go_format_type "github.com/pefish/go-format/type"
	go_test_ "github.com/pefish/go-test"
)

func TestDeepCopy(t *testing.T) {
	results := DeepCopy([]string{"a", "b"})
	go_test_.Equal(t, "a", results[0])
	go_test_.Equal(t, "b", results[1])
}

func TestFormatType_GroupSlice(t *testing.T) {
	strSlice := []string{"a", "b", "c", "d", "e", "f", "g"}
	group := Group(strSlice, &go_format_type.GroupOpts{
		CountPerGroup: 3,
	})
	go_test_.Equal(t, 3, len(group))
	go_test_.Equal(t, 3, len(group[0]))
	go_test_.Equal(t, "a", group[0][0])
	go_test_.Equal(t, "c", group[0][2])
	go_test_.Equal(t, 3, len(group[1]))
	go_test_.Equal(t, "d", group[1][0])
	go_test_.Equal(t, "e", group[1][1])
	go_test_.Equal(t, 1, len(group[2]))
	go_test_.Equal(t, "g", group[2][0])

	group1 := Group(strSlice, &go_format_type.GroupOpts{
		CountPerGroup: 10,
	})
	go_test_.Equal(t, 1, len(group1))
	go_test_.Equal(t, 7, len(group1[0]))
	go_test_.Equal(t, "a", group1[0][0])
	go_test_.Equal(t, "g", group1[0][6])

	group2 := Group(strSlice, &go_format_type.GroupOpts{
		GroupCount: 10,
	})
	fmt.Println(group2)
	// go_test_.Equal(t, 1, len(group1))
	// go_test_.Equal(t, 7, len(group1[0]))
	// go_test_.Equal(t, "a", group1[0][0])
	// go_test_.Equal(t, "g", group1[0][6])
}

func TestFormatClass_SliceToStruct(t *testing.T) {
	type Test struct {
		A uint64 `json:"a"`
		B string `json:"haha"`
	}
	testObj := []Test{}
	err := ToStruct(
		&testObj,
		[]any{
			map[string]any{
				`a`:    100,
				`haha`: `1111`,
			},
			map[string]any{
				`a`:    100,
				`haha`: `1111`,
			},
		},
	)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, true, testObj[0].A == 100)
	go_test_.Equal(t, "1111", testObj[0].B)
}
