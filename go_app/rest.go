package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Name        string `json:Name`
	Description string `json:Desc`
	Number      int    `json:Number`
}

type Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Name: "Test1", Description: "Test description", Number: 5},
		Article{Name: "Test2", Description: "Test description", Number: 3},
		Article{Name: "Test3", Description: "Test description", Number: 7},
		Article{Name: "Test4", Description: "Test description", Number: 9},
	}
	fmt.Println("Endpoint Articles!")
	json.NewEncoder(w).Encode(articles)
}

func handleRequests() {
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
