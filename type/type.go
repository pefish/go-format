package go_format_type

import (
	"strconv"
	"strings"
)

type GroupOpts struct {
	CountPerGroup int
	GroupCount    int // 分成几组，如果组数 a 大于元素个数 b，则结果是《b 组，每组一个元素》
}

type Int64String int64

func (i *Int64String) UnmarshalJSON(b []byte) error {
	// null
	if string(b) == "null" {
		*i = 0
		return nil
	}

	// 去掉引号
	s := strings.Trim(string(b), `"`)
	if s == "" {
		*i = 0
		return nil
	}

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*i = Int64String(v)
	return nil
}

type Float64String float64

func (i *Float64String) UnmarshalJSON(b []byte) error {
	// null
	if string(b) == "null" {
		*i = 0
		return nil
	}

	// 去掉引号
	s := strings.Trim(string(b), `"`)
	if s == "" {
		*i = 0
		return nil
	}

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*i = Float64String(v)
	return nil
}
