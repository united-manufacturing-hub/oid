package oid

var ianaRoot = [6]int{1, 3, 6, 1, 4, 1}
var umhPen = [1]int{59193}

func GetIanaRoot() []int {
	return ianaRoot[:]
}

func GetUMHPen() []int {
	return umhPen[:]
}

func ConcatCopyPreAllocate(slices [][]int) []int {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}
	tmp := make([]int, totalLen)
	var i int
	for _, s := range slices {
		i += copy(tmp[i:], s)
	}
	return tmp
}

func SliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
