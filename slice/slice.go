package go_format_slice

import (
	go_format_int "github.com/pefish/go-format/int"
	go_format_type "github.com/pefish/go-format/type"
)

func DeepCopy[T any](slice []T) []T {
	results := make([]T, len(slice))
	copy(results, slice)
	return results
}

func Group[T any](slice []T, ops *go_format_type.GroupOpts) [][]T {
	resultGroup := make([][]T, 0)

	countPerGroup := ops.CountPerGroup
	if countPerGroup == 0 {
		groupCount := ops.GroupCount
		if groupCount == 0 {
			groupCount = 1
		}
		countPerGroup = len(slice) / ops.GroupCount
		if len(slice)%ops.GroupCount > 0 {
			countPerGroup += 1
		}
	}

	intValues := go_format_int.Group(len(slice), countPerGroup)

	for i, intValue := range intValues {
		start := 0
		if i > 0 {
			start = i * intValues[i-1]
		}
		resultGroup = append(resultGroup, slice[start:start+intValue])
	}
	return resultGroup
}
