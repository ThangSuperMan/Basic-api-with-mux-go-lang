package main

import (
	"BasicApi/structs"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io/ioutil"
	"net/http"
)

var articles []structs.Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("homePage")
	t, _ := template.ParseFiles("./views/index.html")
	t.Execute(w, nil)
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println(vars)
	fmt.Fprintf(w, "Key: %s", key)

	for i := range articles {
		fmt.Println(articles[i])
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(reqBody)
}

func addingArticlePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./views/create-article.html")

	t.Execute(w, nil)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(false)

	myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./assets"))))

	myRouter.HandleFunc("/add_article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article", addingArticlePage).Methods("GET")
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	http.ListenAndServe(":3000", myRouter)
}

func main() {
	articles = []structs.Article{
		structs.Article{Id: "1", Title: "Hello", Desc: "Article description", Content: "Article content"},
		structs.Article{Id: "2", Title: "Hello 2", Desc: "Article description", Content: "Article content"},
	}

	handleRequests()
}
