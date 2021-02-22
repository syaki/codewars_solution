package solution

func Solution(ar []int) int {
	// return smallest possible sum of all numbers in array
	min := findMin(ar)
	res := transform(ar, min)
	for findMin(res) < min {
		min = findMin(res)
		res = transform(ar, min)
	}
	return sum(res)
}

func sum(ar []int) int {
	res := 0
	for _, v := range ar {
		res += v
	}
	return res
}

func findMin(ar []int) int {
	min := ar[0]
	for i := 1; i < len(ar); i++ {
		if min > ar[i] {
			min = ar[i]
		}
	}
	return min
}

func transform(ar []int, min int) []int {
	res := make([]int, len(ar))
	for i := range ar {
		temp := ar[i] % min
		if temp == 0 {
			res[i] = min
		} else {
			res[i] = temp
		}
	}
	return res
}

func CleverSolution(ar []int) int {
	result := ar[len(ar)-1]

	for i := len(ar) - 2; i >= 0; i-- {
		result = gcd(ar[i], result)
	}

	return result * len(ar)
}

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}
