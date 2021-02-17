package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Autor string `json:"autor"`
}

var Books []Book

func apiVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: apiVersion")
	fmt.Fprintf(w, "v0.1")
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getAllBooks")
	json.NewEncoder(w).Encode(Books)
}

func getOneBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getOneBooks")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, book := range Books {
		if book.Id == key {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func createNewBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: createNewBook")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)
	Books = append(Books, book)

	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: updateBook")
	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newBook Book
	json.Unmarshal(reqBody, &newBook)

	for index, book := range Books {
		if book.Id == key {
			Books[index] = newBook
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: deleteBook")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, book := range Books {
		if book.Id == id {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/version", apiVersion)
	myRouter.HandleFunc("/book", getAllBooks).Methods("GET")
	//myRouter.HandleFunc("/book", createNewBook).Methods("POST")
	//myRouter.HandleFunc("/book/{id}", updateBook).Methods("PUT")
	//myRouter.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	myRouter.HandleFunc("/book/{id}", getOneBook)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	Books = []Book{
		{Id: "1", Title: "book 1", Desc: "The first book", Autor: "Unknow"},
		{Id: "2", Title: "book 2", Desc: "The second book", Autor: "Unknow"},
	}

	handleRequest()
}
