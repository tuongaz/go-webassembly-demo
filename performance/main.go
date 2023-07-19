package main

import (
	"syscall/js"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func nthPrime(n int) int {
	count := 0
	num := 2
	for {
		if isPrime(num) {
			count++
			if count == n {
				return num
			}
		}
		num++
	}
}
func main() {
	js.Global().Set("nthPrime", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nthPrime(args[0].Int())
	}))

	<-make(chan bool)
}
