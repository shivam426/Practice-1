
package main
import "fmt"
func main ()  {
	var a int
	fmt.Println("enter a ")
	fmt.Scan(&a)
	var b int
	fmt.Println("enter b ")
	fmt.Scan(&b)
var cal string
fmt.Println("enter ur choice:")
fmt.Scan(&cal)
switch cal {
case "+":
	fmt.Println(a+b)
	break
case "-":
	fmt.Println(a-b)
	break
case "*":
	fmt.Println(a*b)
	break
case "/":
	fmt.Println(a/b)
	break
default:
	fmt.Println("invalid data")
}
}