package binarysearch

func SearchInts(arr []int, key int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := (l + r) / 2
		switch {
		case arr[m] == key:
			return m
		case arr[m] > key:
			r = m - 1
		case arr[m] < key:
			l = m + 1
		}
	}
	return -1
}
