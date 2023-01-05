package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	if !IsPrime(n) {
		fmt.Println("Bukan Prima")
	} else {
		fmt.Println("Prima")
	}
}

func IsPrime(n int) bool {
	var prime bool
	var i int
	prime = true
	i = 2
	if n == 1 {
		prime = false
	}
	for prime == true && i < n {
		if n%i == 0 {
			prime = false
		} else {
			i++
		}
	}
	return prime
}
