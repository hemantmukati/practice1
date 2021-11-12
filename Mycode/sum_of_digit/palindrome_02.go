package main

import "fmt"

func main() {
	x := "1332331"

	flag := palindrom(x)

	if flag {
		fmt.Println("Palindrom number")
	} else {
		fmt.Println("No Palindrom number")
	}
}

func palindrom(x string) bool {
	flag := true
	for i := 0; i < len(x); i++ {
		if x[i] != x[len(x)-1-i] {
			return false
		}
	}
	return flag
}
