package main

import "fmt"

func main() {

	arr := []int{4, 12, 5, 21, 51, 14, 6, 49, 55}
	largestNum := 4
	for i := 0; i < len(arr); i++ {

		largestNum = arr[0]

		if largestNum < arr[i] {
			largestNum = arr[i]
		}
	}
	fmt.Println(largestNum)
}
