package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kylemlee/klapp/src/driver"
)

func HomeHanlder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {

	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)

	if err != nil {
		fmt.Println("an error occured")
		panic(err.Error())
	}

	defer connection.SQL.Close()

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", HomeHanlder)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":80", r))

}
