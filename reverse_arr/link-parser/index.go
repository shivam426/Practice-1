package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

type Data struct {
	Href string
	Link string
}

// ReadHTML for reading from the dot html file
func ReadHTML(filename string) string {
	spt, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	str := string(spt)
	myreader := strings.NewReader(str)
	Parser(myreader)
	return str
}

// Parser for parsing the string
func Parser(r *strings.Reader) []Data {

	t, err := html.Parse(r)
	if err != nil {
		fmt.Println(err)
	}
	Ds(t, "")
	return nil
}

// Ds for creating the nodes
func Ds(n *html.Node, pad string) string {
	fmt.Println(pad, n.Data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Ds(c, pad+" ")
	}
	return n.Data
}

func main() {
	fileName := "exp.html"
	ReadHTML(fileName)

}
