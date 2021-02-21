package solution

func Decode(roman string) int {
	codeMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	romanLen := len(roman)
	for i := 0; i < romanLen-1; i++ {
		curCode := roman[i]
		nextCode := roman[i+1]
		if codeMap[curCode] < codeMap[nextCode] {
			res -= codeMap[curCode]
		} else {
			res += codeMap[curCode]
		}
	}
	res += codeMap[roman[romanLen-1]]
	return res
}
