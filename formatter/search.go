package formatter

import "sort"

func binarySearch(s []int, n int) bool {
	idx := sort.SearchInts(s, n)

	return idx < len(s) && s[idx] == n
}
