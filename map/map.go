package go_format_map

import (
	go_format_slice "github.com/pefish/go-format/slice"
	go_format_type "github.com/pefish/go-format/type"
)

// 生成的结果中，key 的排列是无序的
func Group[F comparable, T any](map_ map[F]T, ops *go_format_type.GroupOpts) []map[F]T {
	resultGroup := make([]map[F]T, 0)

	keys := make([]F, 0, len(map_))
	for key := range map_ {
		keys = append(keys, key)
	}

	keysByGroup := go_format_slice.Group(keys, ops)
	for _, keysInGroup := range keysByGroup {
		groupMap := make(map[F]T)
		for _, keyInGroup := range keysInGroup {
			groupMap[keyInGroup] = map_[keyInGroup]
		}
		resultGroup = append(resultGroup, groupMap)
	}

	return resultGroup
}
