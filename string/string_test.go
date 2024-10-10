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
