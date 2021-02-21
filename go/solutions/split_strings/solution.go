package solution

import "regexp"

func Solution(str string) []string {
	res := []string{}

	for idx, ch := range str {
		if idx%2 == 0 {
			res = append(res, string(ch))
		} else {
			temp := res[len(res)-1]
			res[len(res)-1] = temp + string(ch)
		}
	}

	if len(res[len(res)-1]) < 2 {
		temp := res[len(res)-1]
		res[len(res)-1] = temp + "_"
	}

	return res
}

func CleverSolution(str string) []string {
	return regexp.MustCompile(".{2}").FindAllString(str+"_", -1)
}
