package main

import "fmt"

func main() {
	var a int = 4
	x := factorial(a)
	fmt.Println(x)
	c,d:=factSum(x)
	fmt.Println(c,d)
}
func factorial(f int) int {
	fact := 1
	for i := 1; i <= f; i++ {
		fact = fact * i
	}
	return fact
}
func factSum(x int) (int,int) {

	sum := 0
	for x > 0 {
		sum = sum + (x % 10)
		x = x / 10
	}
    	return sum,x
}
