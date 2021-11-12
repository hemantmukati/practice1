package main

import "fmt"

func main(){

       arr:= [] int {1,2,0,4,7,3,8,9,12,6}
    
      fmt.Println("before sorting:" , arr)
      for i:=0; i<=len(arr)-1; i ++{

      for j:=i+1; j<=len(arr)-1;j++{

      	if arr[i] > arr[j] {
           
          temp:= arr[j]

           arr[j]= arr[i]	

           arr[i] = temp
      	}
      }
      } 
     fmt.Println("after sorting:",arr)
}

