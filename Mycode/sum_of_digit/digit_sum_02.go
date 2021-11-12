package main

import "fmt"

func main() {
	var num = 99999999
	x := 10

	sum := 0
	for num > 0 {

		sum += num % x
		num = num / x

	}

	fmt.Println(sum)
}
