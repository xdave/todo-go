package sqldb

import (
	"database/sql"
	"fmt"
	"os"
	"todo-go/structs"

	_ "github.com/lib/pq"
)

var db *sql.DB

var (
	host     = os.Getenv("DB_HOST")     //"127.0.0.1"
	port     = os.Getenv("DB_PORT")     // "5432"
	user     = os.Getenv("DB_USERNAME") // "root"
	password = os.Getenv("DB_PASSWORD") // "secret"
	dbname   = "root"
)

func OpenConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	mydb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = mydb.Ping()
	if err != nil {
		panic(err)
	}
	db = mydb
	createTable()
}

func createTable() {
	stmt := `
	CREATE TABLE todo(id INT GENERATED ALWAYS AS IDENTITY,
	 ip_address VARCHAR(200) NOT NULL, text VARCHAR(2000) NOT NULL,
	 checked BOOLEAN NOT NULL, created_at timestamptz NOT NULL DEFAULT NOW());`

	_, err := db.Exec(stmt)
	if err != nil {
		fmt.Println(err)
	}
}

func ShowTables() {
	rows, err := db.Query(`SELECT * FROM pg_catalog.pg_tables
	WHERE schemaname != 'pg_catalog' AND 
		schemaname != 'information_schema';`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	cols, _ := rows.Columns()
	// fmt.Println(rows.Columns())
	for i := range cols {
		fmt.Println(cols[i])
	}
	fmt.Println(rows.Next())

}

func Close() {
	db.Close()
}

func InsertTodo(t *structs.TodoInsert) (int64, error) {
	stmnt := `INSERT INTO todo (ip_address, text, checked) VALUES ($1, $2, $3) RETURNING id`
	var id int64
	err := db.QueryRow(stmnt, t.Ip_Address, t.Text, t.Checked).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Error adding todo: %v", err)
	}
	// fmt.Printf("%v %[1]T ",id)
	return id, err
}

func GetTodos() ([]structs.TodoQuery, error) {
	rows, err := db.Query(`SELECT id, text, checked FROM todo;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var ts []structs.TodoQuery
	for rows.Next() {
		var t structs.TodoQuery
		if err := rows.Scan(&t.Id, &t.Text, &t.Checked); err != nil {
			return ts, err
		}
		ts = append(ts, t)
	}
	if err = rows.Err(); err != nil {
		return ts, err
	}
	return ts, nil

}

func UpdateTodoChecked(id int64, check bool) error {
	stmnt := `UPDATE todo SET checked=($1) where id=($2)`
	_, err := db.Exec(stmnt, check, id)
	if err != nil {
		return fmt.Errorf("Error toggling todo: %v", err)
	}
	return err
}

func DeleteTodo(id int64) error {
	stmnt := `DELETE FROM todo where id=($1)`
	_, err := db.Exec(stmnt, id)
	if err != nil {
		return fmt.Errorf("Error deleting todo: %v", err)
	}
	return err
}
