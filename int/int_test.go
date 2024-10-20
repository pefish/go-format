package go_format_int

import (
	"fmt"
	"testing"
)

func TestGroupInt(t *testing.T) {
	results := Group(34, 10)
	fmt.Println(results)
}
