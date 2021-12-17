package f2

import "math"

func InSliceString(array_val string, array []string) bool {
	for _, v := range array {
		if v == array_val {
			return true
		}
	}
	return false
}

func SliceChunkString(s []string, size int) [][]string {
	if size < 1 {
		return [][]string{}
	}
	length := len(s)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]string
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, s[i*size:end])
		i++
	}
	return n
}
