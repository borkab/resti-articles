package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type article struct {
	Title   string `json:"title"`
	Disc    string `json:"disc"`
	Content string `json:"content"`
}

type Articles []article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := article{
		"Test Title",
		"Test Description",
		"Hello World",
	}

	fmt.Println("Endpoint Hit: all Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

func handleRequest() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/article", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
