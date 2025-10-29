package go_format_string

import (
	"fmt"
	"testing"

	go_format_type "github.com/pefish/go-format/type"
	go_test_ "github.com/pefish/go-test"
)

func TestBetweenAnd(t *testing.T) {
	results := BetweenAnd("rgw4tq874hrwuh8gw4rg89iju8", "4", "8")
	go_test_.Equal(t, 3, len(results))
	go_test_.Equal(t, "tq", results[0])
	go_test_.Equal(t, "hrwuh", results[1])
	go_test_.Equal(t, "rg", results[2])
}

func TestInsert(t *testing.T) {
	result, err := Insert("9", "012345", 6)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "0123459", result)
}

func TestIndexes(t *testing.T) {
	results := Indexes("59852645762485752776543352762", "762")
	go_test_.Equal(t, 8, results[0])
	go_test_.Equal(t, 26, results[1])
}

func TestLastIndex(t *testing.T) {
	result := LastIndex("123456789", []string{"8", "6"})
	go_test_.Equal(t, 7, result)
	result = LastIndex("123845679", []string{"8", "6"})
	go_test_.Equal(t, 6, result)
}

func TestGroupString(t *testing.T) {
	str := "gsd449998d88gsgsrt"
	results := Group(str, &go_format_type.GroupOpts{
		GroupCount: 3,
	})
	fmt.Println(results)
}

func TestTrimPunct(t *testing.T) {
	str := "（混合Symbols!）"
	result := TrimPunct(str)
	fmt.Println(result)
}
