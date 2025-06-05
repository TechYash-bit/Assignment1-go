package main
//largest of three
import "fmt"
func main() {
 var a,b,c int 
 fmt.Println("Enter the 1 number")
 fmt.Scan(&a)
 fmt.Println("Enter the 2 number")
 fmt.Scan(&b)
 fmt.Println("Enter the 3 number")
 fmt.Scan(&c)

 if a>b && a>c {
	fmt.Print("largest is ",a)
 }else if b>c{
	fmt.Println(" largest is ",b)
 }else{
	fmt.Println("largest is ",c)
 }


}