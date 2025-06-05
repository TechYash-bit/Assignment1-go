package main

//swap two number
import "fmt"
func main() {
	fmt.Println("enter any number")
	var a int
	fmt.Scan(&a)
	fmt.Println("enter the other number")
	var b int
	fmt.Scan(&b)

	fmt.Println("before swaping  a :",a," b : ",b )

	var c int
	c=a
	a=b
	b=c
	
	fmt.Println("After swap  a: ",a," b: ",b )

}