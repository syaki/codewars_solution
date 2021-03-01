package solution

// timed out
func ProductSum(a []int, m int) int {
	// your code here
	var cache = make(map[int]map[int][]int)
	res := 0

	var perm func([]int, int, int) []int
	perm = func(a []int, i, m int) []int {
		if 0 == m {
			return []int{1}
		}

		res := []int{}
		for ; i <= len(a)-m; i++ {
			_, ok := cache[m-1]
			if !ok {
				cache[m-1] = make(map[int][]int)
			}
			temp, ok := cache[m-1][i+1]
			if !ok {
				temp := perm(a, i+1, m-1)
				cache[m-1][i+1] = temp
			}
			for _, v := range temp {
				res = append(res, v*a[i])
			}
		}
		return res
	}
	temp := perm(a, 0, m)
	for _, v := range temp {
		res += v
	}
	return res
}

func CleverProductSum(a []int, m int) int {
	const D = 1e9 + 7
	r := make([]int, m+1)
	r[0] = 1
	for i := 0; i < len(a); i++ {
		for j := m; j > 0; j-- {
			r[j] = (r[j] + r[j-1]*a[i]) % D
		}
	}
	return r[m]
}
