package main

import "fmt"

func main() {
	x := "13342331"

	flag := palindrom(x)

	if flag {
		fmt.Println("Palindrom number")
	} else {
		fmt.Println("No Palindrom number")
	}
}

func palindrom(x string) bool {
	for i := 0; i < len(x)/2; i++ {
		if x[i] != x[len(x)-1-i] {
			return false
		}
	}
	return true
}
