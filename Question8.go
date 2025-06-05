package main
// find the given numbre is positive or not
import "fmt"
func main() {
	var a int 
	fmt.Println("enter any random number ")
	fmt.Scan(&a)

	if a==0{
		fmt.Println("the number is zero")
	}else if a>0{
		fmt.Println("the number is positive")
	}else{
		fmt.Println("the number is negative")
	}


}