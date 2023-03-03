package main

import ("fmt")

func main()  {
	y:="I want know a best way to check if some char is vowel"
	for _, v:= range y{
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' ||v == 'A' ||v == 'E' ||v == 'O' ||v == 'I' ||v == 'O' ||v == 'U' {
			v++

		}
	fmt.Printf("%s",string(v))	
	}
fmt.Println(" ")	
}