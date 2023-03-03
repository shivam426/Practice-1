package main

import "fmt"

// Person for the info
type Person struct {
	Name    string
	Age     int8
	Address string
}

func (p *Person) String() string {
	return fmt.Sprint(p.Name, p.Age, p.Address)
}
func main() {
	var p Person
	p = Person{"shivam", 24, "Puspa niwas 1446/4096 santoshi vihar,kanan vihar phase 2,patia "}
	fmt.Println(p)
}
