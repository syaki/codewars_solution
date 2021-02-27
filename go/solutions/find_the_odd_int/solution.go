package solution

func FindOdd(seq []int) int {
	// your code here
	timesMap := map[int]int{}
	for _, v := range seq {
		timesMap[v] += 1
	}
	for k, v := range timesMap {
		if v%2 != 0 {
			return k
		}
	}
	return 0
}

func CleverFindOdd(seq []int) int {
	res := 0
	for _, x := range seq {
		res ^= x
	}
	return res
}
