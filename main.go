package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Fname string `json:"firstname"`
	Lname string `json:"lastname"`
}

var Books = []Book{{1, "Dark Night", Author{"Vikash", "Parashar"}}, {2, "Dark Dreams", Author{"Pawan", "Bhardwaj"}}, {3, "Star Light", Author{"Romeo", "Siwach"}}}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", AllBooks).Methods("GET") // successful
	r.HandleFunc("/api/books/{id}", SingleBook).Methods("GET")
	r.HandleFunc("/api/books", NewBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", DeleteBook).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliation/json")

	if err := json.NewEncoder(w).Encode(Books); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func SingleBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliation/json")
	params := mux.Vars(r)
	id := params["id"]
	iid, err := strconv.Atoi(id)
	if err != nil {
		log.Println("can not convert string type Id into int")
	}

	for _, v := range Books {
		if v.Id == iid {
			json.NewEncoder(w).Encode(v)
		} else {
			w.WriteHeader(http.StatusNotFound)
			continue
		}
	}

}

func NewBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliation/json")
	var b []Book
	json.NewDecoder(r.Body).Decode(&b)
	Books = append(Books, b[0:]...)
	if err := json.NewEncoder(w).Encode(Books); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	// 1st way to do this

	w.Header().Set("Content-Type", "appliation/json")
	var book Book
	params := mux.Vars(r)
	id := params["id"]
	iid, err := strconv.Atoi(id)
	if err != nil {
		log.Println("can not convert string type Id into int")
	}

	json.NewDecoder(r.Body).Decode(&book)
	book.Id = iid
	for i, v := range Books {
		if v.Id == iid {
			Books = append(Books[:i], Books[i+1:]...)
			// json.NewEncoder(w).Encode(Books)
		} else {
			continue
		}
	}

	Books = append(Books, book)
	json.NewEncoder(w).Encode(Books)

	// 2ne way to do that

	// w.Header().Set("Content-Type", "appliation/json")
	// var book Book
	// params := mux.Vars(r)
	// id := params["id"]
	// iid, err := strconv.Atoi(id)
	// if err != nil {
	// 	log.Println("can not convert string type Id into int")
	// }

	// book.Id = iid
	// for _, v := range Books {
	// 	if v.Id == iid {
	// 		json.NewDecoder(r.Body).Decode(&book)
	// 		v.Title = book.Title
	// 		v.Author = book.Author
	// 		fmt.Println(v.Title)
	// 		fmt.Println(v.Author)
	// 		// Books = append(Books[:i], Books[i+1:]...)
	// 		// json.NewEncoder(w).Encode(Books)

	// 		// Books = append(Books, book)
	// 		// json.NewEncoder(w).Encode(Books)
	// 	} else {
	// 		continue
	// 	}
	// }

	// Books = append(Books, book)
	// json.NewEncoder(w).Encode(Books)
	// fmt.Println(book)
	// fmt.Println(Books)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliation/json")

	params := mux.Vars(r)
	id := params["id"]
	iid, err := strconv.Atoi(id)
	if err != nil {
		log.Println("can not convert string type Id into int")
	}

	for i, v := range Books {
		if v.Id == iid {
			Books = append(Books[:i], Books[i+1:]...)
			json.NewEncoder(w).Encode(Books)
		} else {
			continue
		}
	}
}
