package main

func Fibo(n int) int {
	if n <= 1 {
		return n
	}
	return Fibo(n-1) + Fibo(n-2)
}
