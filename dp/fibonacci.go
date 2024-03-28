package dp

func Fibonacci(n int) int {
	if n < 0 {
		return -1
	}
	back2 := 0
	if n == 0 {
		return back2
	}
	back1 := 1
	if n == 1 {
		return back1
	}

	for i := 2; i <= n; i++ {
		back1, back2 = back1+back2, back1
	}

	return back1
}
