package main

import (
	"fmt"
	"kanolibrary/routes"
	"net/http"
)

type books struct {
	Name      string `bson:"name"`
	Author    string `bson: "author"`
	Publisher string `bson: "publisher"`
	Synopsis  string `bson: "Synopsis"`
	Page      string `bson : "page"`
}

func main() {
	// crud.Insert("books", books {Name: "Mantappu", Author : "Jerome Polin"})
	//crud.Insert("books", books {Name: "Nebula", Author : "Tere Liye", Publisher: "Erlangga", Synopsis: "kisah hidup mahasiswa jepang", Page: "300"})
	// crud.Find("books", books {})
	// crud.Update()
	// crud.Remove()
	http.HandleFunc("/books", routes.FindAllBooks)
	http.HandleFunc("/books/find", routes.FindByIdBook)
	http.HandleFunc("/books/insert", routes.InsertBook)
	http.HandleFunc("/books/update", routes.UpdateBook)
	http.HandleFunc("/books/delete", routes.DeleteBook)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil) 
}
