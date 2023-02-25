package main

import (
	"fmt"
	"log"
	"net/http"
	"acron/animation"
)

func main() {
	animation.Animation()
	http.HandleFunc("/", homePage)
	log.Println("Starting web server on port 8080")
	animation.Animation()
	http.ListenAndServe(":8080", nil)


}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Sina")
}