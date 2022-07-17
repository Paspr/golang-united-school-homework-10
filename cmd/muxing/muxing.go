package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{param}", paramHandler).Methods("GET")
	router.HandleFunc("/bad", badHandler).Methods("GET")
	router.HandleFunc("/data", dataHandler).Methods("POST")
	router.HandleFunc("/headers", headersParam).Methods("POST")
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func paramHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, "+vars["param"]+"!")
}

func badHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Internal server error")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	b, error := io.ReadAll(r.Body)
	if error == nil {
		fmt.Fprintf(w, "I got message:\n"+string(b))
	}
}

func headersParam(w http.ResponseWriter, r *http.Request) {

	headerA := r.Header.Get("a")
	headerB := r.Header.Get("b")

	valueA, errA := strconv.Atoi(headerA)
	valueB, errB := strconv.Atoi(headerB)

	if errA == nil && errB == nil {
		w.Header().Set("a+b", strconv.Itoa(valueA+valueB))
	}
}
