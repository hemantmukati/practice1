package main

import "fmt"
 
func main (){ 
            x:=[] int {1,2,3,4,5}
			
			fmt.Println("before reverse:", x)

			 reverseArr:= make([]int,len(x))
			 
			 
			 for i:=0; i<len(x); i++{
			 

			 reverseArr[len(x)-1-i] = x[i]

             
		      
		}
          fmt.Println("after reverse:", reverseArr)
   } 