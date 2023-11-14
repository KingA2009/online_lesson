package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test title", Desc: "Test description", Content: "Hello world "},
	}
	fmt.Println("Endpoint Hit: All articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint Hit")
}

func handlerRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main() {
	handlerRequests()
}
