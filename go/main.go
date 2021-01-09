package main 

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func helloMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello")
}



func main(){
	//Router
	r:=mux.NewRouter()
	r.HandleFunc("/firstTry", helloMessage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}