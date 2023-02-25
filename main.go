package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)


}


func homePage(w http.ResponseWriter, r *http.Request) {
	html := `<strong>Hello Sina</strong>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}