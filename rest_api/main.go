package main
import (
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", returnAllUsers).Methods("GET")
	router.HandleFunc("/users/add", insertUsers).Methods("POST")
	router.HandleFunc("/users/update", updateUsers).Methods("PUT")
	router.HandleFunc("/users/delete", deleteUsers).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connection port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}