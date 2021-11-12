package main
import "fmt"

func main(){
      
	 x:=factorial()
	 fmt.Println(x)
	 
    }
    func factorial() int{  	
	 fact:=1    
	for i:=1;i<=5;i++{
	   fact = fact * i   
  }	
	return fact
 
}