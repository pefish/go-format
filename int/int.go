package go_format_int

// 对数值进行分组。例如 35 使用 10 分组结果是 [10,10,10,5]
func Group[T int | uint | int64 | uint64](number T, sliceBy T) []T {
	results := make([]T, 0)
	var start, end T = 0, 0
	for {
		start = end
		end += sliceBy
		if end > number {
			end = number
		}
		results = append(results, end-start)
		if end >= number {
			break
		}
	}
	return results
}
