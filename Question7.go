package main
//find the given number is even or oddd
import "fmt"
func main() {
	fmt.Println("Enter the 1 number")
	var a int
	fmt.Scan(&a)
	 if a%2==0{
		fmt.Println("the given number is even number ")
	 }else{
		fmt.Println("the given number is odd number ")
	 }


}