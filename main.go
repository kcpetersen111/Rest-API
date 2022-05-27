package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct{
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title:"Test",
				Desc:"Test",
				Content:"Test"},
	}
	log.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)

}

func testPostArticles(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Test post endpoint\n")
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hit the homepage\n")
}

func handleRequest(){

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/Articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/Articles", testPostArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080",myRouter))
}

func main() {
	handleRequest()
}


http://localhost:8080/Articles
