package piscine

func Max(arr []int) int {
	k := 0
	ok := true
	for _, a := range arr {
		if a > k {
			k = a
		}
		if ok {
			k = a
			ok = false
		}
	}
	return k
}
