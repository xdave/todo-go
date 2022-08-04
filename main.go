package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

type Todo struct {
	Id   uuid.UUID
	Text string
	Done bool
}
type updated_todo struct {
	Id   uuid.UUID
	Done bool
}

var mux = http.NewServeMux()

var id1, id2 = uuid.New(), uuid.New()
var db = map[uuid.UUID]Todo{
	id1: {Id: id1, Text: "Wake up", Done: true},
	id2: {Id: id2, Text: "Do 10 backflips", Done: false},
}

func list_todo(w http.ResponseWriter, r *http.Request) {
	jsonValue, _ := json.Marshal(db)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err := json.NewEncoder(w).Encode(jsonValue)
	if err != nil {
		fmt.Printf("error can't list todo \n %v \n", err)
		panic("aaaaah errored out!")

	}
}

func add_todo(w http.ResponseWriter, r *http.Request) {
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	i := uuid.New()
	db[i] = Todo{Id: i, Done: false, Text: string(resp)}
	fmt.Println(db)
	// w.Write([]byte(i.String()))
	io.WriteString(w, i.String())
}

func update_checked(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var ut updated_todo
	er := decoder.Decode(&ut)
	if er != nil {
		fmt.Println(er)
		panic("aaaaah errored out!")
	}
	var oldTodo Todo = db[ut.Id]
	// fmt.Printf("current db entry %v, received ut: %v\n \n", oldTodo, ut)
	oldTodo.Done = ut.Done
	// fmt.Printf("updated db entry %v @ %p \n ", oldTodo, &oldTodo)
	db[ut.Id] = oldTodo
	// fmt.Println(db[ut.Id])
	// fmt.Println("---")

}

func main() {

	mux.Handle("/", http.FileServer(http.Dir("./frontend")))
	mux.HandleFunc("/add_todo", add_todo)
	mux.HandleFunc("/check_todo", update_checked)
	mux.HandleFunc("/list", list_todo)
	mux.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "OK") })
	var port string = ":8000"
	fmt.Printf("starting server at port %s\n", port)
	fmt.Println(http.ListenAndServe(port, mux))

}
