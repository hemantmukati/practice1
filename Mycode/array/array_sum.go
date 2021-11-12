package main 

import "fmt"

func main112 (){

      arr := []int {5,8,2,4,1,9,12,23,24,25,45,56,67}
       sum:=0
     for i:=0; i< len(arr);i ++{
         
       
         sum = sum + arr[i]

     
        fmt.Println(sum) 
    }

}