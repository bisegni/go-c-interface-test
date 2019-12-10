package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	guuid "github.com/google/uuid"

	"github.com/bisegni/go-c-interface-test/adapter"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

type process struct {
	ID     string                 `json:"id"`
	Config adapter.ExternalSource `json:"source"`
}

func starProcess(w http.ResponseWriter, r *http.Request) {
	var newProc process
	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))
	if err != nil {
		fmt.Fprintf(w, "Insert data for the process to start")
	}
	// Unmarshal the process from json
	json.Unmarshal(reqBody, &newProc)

	//create new id that will became the folder for the processing data
	newProc.ID = guuid.New().String()
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newProc)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/process", starProcess).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
