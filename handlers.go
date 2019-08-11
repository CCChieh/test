package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Welcome!")
}
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(todos); err != nil {
		panic(err)
	}
}
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	_, _ = fmt.Fprintln(w, "Todo show:", todoId)
}
