package main
import "fmt"
//addition of two number
func main() {
	var a,b int

	fmt.Println("Enter the first number")
	fmt.Scan(&a)
	fmt.Println("Enter the second number")
	fmt.Scan(&b)

	fmt.Println("the sum of ",a,"  and ", b," is",a+b)

}