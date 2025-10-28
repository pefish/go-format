package go_format_map

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	go_format_type "github.com/pefish/go-format/type"
)

func TestGroup(t *testing.T) {
	strSlice := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
		"d": "4",
		"e": "5",
		"f": "6",
		"g": "7",
	}
	group := Group(strSlice, &go_format_type.GroupOpts{
		CountPerGroup: 3,
	})
	spew.Dump(group)
}
