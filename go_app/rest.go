package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Articule struct {
	Name        string `json:Name`
	Description string `json:Desc`
	Number      int    `json:Number`
}

type Articules []Articule

func returnAllArticules(w http.ResponseWriter, r *http.Request) {
	articules := Articules{
		Articule{Name: "Test", Description: "bgfbgfbgf", Number: 2},
		Articule{Name: "Test2", Description: "blalvFSD", Number: 3},
	}
	fmt.Println("Endpoint Articules!")
	json.NewEncoder(w).Encode(articules)
}

func handleRequests() {
	http.HandleFunc("/articules", returnAllArticules)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
