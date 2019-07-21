package exp

func Exp(a int, b int) int {
	ans := a
	for i := 1; i<b; i++ {
		ans *= a
	}
	return ans
}