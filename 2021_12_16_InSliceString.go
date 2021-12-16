package f2

func InSliceString(array_val string, array []string) bool {
	for _, v := range array {
		if v == array_val {
			return true
		}
	}
	return false
}
