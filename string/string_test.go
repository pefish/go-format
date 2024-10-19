package go_format_string

import (
	"testing"

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
	result = LastIndex("12345679", []string{"8", "6"})
	go_test_.Equal(t, 5, result)
}
