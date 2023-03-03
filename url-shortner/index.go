package main

import (
	"fmt"
	"net/http"
)

func main() {
	Ptu := map[string]string{
		"/go":   "https://go.dev/tour/welcome/1",
		"/news": "https://www.indiatoday.in/news.html",
	}
	route := http.NewServeMux()
	handler := MapHandler(Ptu, route)
	// 	yaml := `
	// - path: /urlshort
	//   url: https://go.dev/tour/welcome/1
	// - path: /urlshort-final
	//   url: https://www.indiatoday.in/news.html
	// `

	route.HandleFunc("/", Hello)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)

}

// Hello for testing
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World !")
}

// MapHandler for the string
func MapHandler(Ptu map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		if d, ok := Ptu[p]; ok {
			http.Redirect(w, r, d, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler for yaml
// func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
// 	return nil, nil
// }
