package main
//make a table of any number
import "fmt"
func main() {
	fmt.Println("enter any kind of number for which you have to create a table")
	var a,ans int
	fmt.Scan(&a)
	ans=1
	fmt.Println("the table of ",a,"is given below")

	for i:=1;i<11;i++{
		ans=a*i
		fmt.Println(a,"*",i,"=",ans)
	}


}