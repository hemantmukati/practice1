package main 

import "fmt"

func main (){

    arr:= [] int {3,6,1,7,12,0,9,10,4}
    fmt.Println("before sorting:", arr)

   for i:=0; i<len(arr);i++{
   for j:= i+1; j< len(arr); j++{
    if arr[i] > arr[j]{

        temp := arr[j]
        arr[j] = arr[i]
        arr[i] = temp
    }
  }
}
    fmt.Println("after sorting:",arr)
}
 
   
