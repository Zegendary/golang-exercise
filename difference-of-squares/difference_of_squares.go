package diffsquares

func SquareOfSum(n int) (r int) {
	for i := 1; i < n+1; i++ {
		r = r + i
	}
	return r * r
}

func SumOfSquares(n int) (r int) {
	for i := 1; i < n+1; i++ {
		r = r + i*i
	}
	return r
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
