package go_format_type

type GroupOpts struct {
	CountPerGroup int
	GroupCount    int // 分成几组，如果组数 a 大于元素个数 b，则结果是《b 组，每组一个元素》
}
