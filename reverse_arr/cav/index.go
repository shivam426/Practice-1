package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type Story map[string]Chapter

var tpl *template.Template

type handler struct {
	s Story
}

// Chapter for the venture
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func main() {

	j, err := os.Open("gophers.json")
	if err != nil {
		fmt.Println(err)
	}
	d := json.NewDecoder(j)
	var story Story
	if err := d.Decode(&story); err != nil {
		fmt.Println(err)
	}
	h := NewHandler(story)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Fatal(err)
	}

}

// HandlerTpl  for the html template
var HandlerTpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Choose A Venture</title>
</head>
<body>
   <h1>{{.Title}}</h1> 
   {{range .Paragraphs}}
   <p>{{.}}</p>
   {{end}}
<ul>
       {{range .Options}}
       <li><a href="/{{.Arc}}">{{.Text}}</a></li>
       {{end}}
   </ul>
</body>
</html>`

// NewHandler for handler
func NewHandler(s Story) http.Handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.New("").Parse(HandlerTpl))
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]
	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Something went Wrong", http.StatusAccepted)
		}
		return
	}

}
