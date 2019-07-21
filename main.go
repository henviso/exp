package exp

import "fmt"

func Exp(a int, b int) int {
	ans := a
	for i := 0; i<b-1; i++ {
		ans *= a
	}
	return ans
}