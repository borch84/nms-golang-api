
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"./controllers"
	"./driver"
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {

	db = driver.ConnectDB()
	controller := controllers.Controller{}
	router := mux.NewRouter()

	router.HandleFunc("/delete/{server}", controller.RemoveServer(db)).Methods("GET")


	fmt.Println("Server is running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
