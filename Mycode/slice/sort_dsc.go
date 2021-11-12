package main

import "fmt"

func main (){
    
    arr := []int {2,5,3,9,0,12,4,13}

   fmt.Println("before sorting:",arr)
    for i:=0 ; i<len(arr); i++{

    for j:=i+1; j<len(arr);j++{

      if arr[i] < arr[j] {
       
        temp:= arr[i]
        arr[i]= arr[j]
        arr[j]= temp

      }
    }  
    }
    fmt.Println("after sorting:",arr)
}