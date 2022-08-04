package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	Id   uint
	Text string
	Done bool
}

var mux = http.NewServeMux()
var i uint = 2
var db = map[uint]Todo{
	1: {Id: 1, Text: "Wake up", Done: true},
	2: {Id: 2, Text: "Do 10 backflips", Done: false},
}

func list_todo(w http.ResponseWriter, r *http.Request) {
	jsonValue, _ := json.Marshal(db)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err := json.NewEncoder(w).Encode(jsonValue)
	if err != nil {
		fmt.Printf("error can't list todo \n %v \n", err)
	}
}

func add_todo(w http.ResponseWriter, r *http.Request) {
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	i++
	db[i] = Todo{Id: i, Done: false, Text: string(resp)}
	fmt.Println(db)
}

func update_checked(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {

	mux.Handle("/", http.FileServer(http.Dir("./frontend")))
	mux.HandleFunc("/add_todo", add_todo)
	mux.HandleFunc("/list", list_todo)
	mux.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "OK") })
	var port string = ":8000"
	fmt.Printf("starting server at port %s\n", port)
	fmt.Println(http.ListenAndServe(port, mux))

}
