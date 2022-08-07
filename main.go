package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"todo-go/sqldb"
	"todo-go/structs"
)

var mux = http.NewServeMux()

func handleError(w http.ResponseWriter, err error){
	fmt.Printf("an error has occured \n %v \n", err)
	http.Error(w, "oops", 400)
}
func list_todo(w http.ResponseWriter, r *http.Request) {

	ts, err := sqldb.GetTodos()
	if err != nil {
		handleError(w, err)
		return
	}
	fmt.Println(ts)
	jsonValue, _ := json.Marshal(ts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(jsonValue)
	if err != nil {
		handleError(w, err)
		return
	}
}

func add_todo(w http.ResponseWriter, r *http.Request) {
	origin := r.Header["Origin"][0]
	// fmt.Println(origin)
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, err)
		return
	}
	id, err := sqldb.InsertTodo( &structs.TodoInsert{Checked: false, Text: string(resp), Ip_Address: origin})
	if err != nil {
		handleError(w, err)
		return
	}
	io.WriteString(w, fmt.Sprintf("%d", id))
}

func update_checked(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var ut structs.TodoCheck
	er := decoder.Decode(&ut)
	if er != nil {
		handleError(w, er)
		return
	}
	er = sqldb.UpdateTodoChecked(ut.Id, ut.Checked)
	if er != nil {
		handleError(w, er)
		return
	}
}
func delete_todo(w http.ResponseWriter, r *http.Request) {
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, err)
		return
	}
	s := string (resp)
	id, err  := strconv.ParseInt(s, 10, 64)
	if err != nil {
		handleError(w, err)
		return
	}
	err = sqldb.DeleteTodo(id)
	if err != nil {
		handleError(w, err)
		return
	}
}


func main() {

	sqldb.OpenConnection()
	mux.Handle("/", http.FileServer(http.Dir("./frontend")))
	mux.HandleFunc("/add_todo", add_todo)
	mux.HandleFunc("/check_todo", update_checked)
	mux.HandleFunc("/delete_todo", delete_todo)
	mux.HandleFunc("/list", list_todo)
	mux.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "OK") })
	var port string = ":8000"
	fmt.Printf("starting server at port %s\n", port)
	fmt.Println(http.ListenAndServe(port, mux))
	
	defer sqldb.Close()

}
