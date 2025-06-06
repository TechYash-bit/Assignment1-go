package main
//factorial of a number
import "fmt"
func main() {
	fmt.Println("Enter number which factorial is to be calculate ")
	var a,fact int
	fact=1
	fmt.Scan(&a)

	for i:=1;i<=a;i++{
		fact=fact*i
	}
fmt.Println("factorial of ",a,"is ",fact)
}