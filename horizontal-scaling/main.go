package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type fibonacci struct{}

func (f fibonacci) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	positionStr := params["value"]
	position, err := strconv.Atoi(positionStr)
	if err != nil {
		http.Error(w, "invalid value", 400)
		return
	}
	value := f.getValue(position)
	w.Write([]byte(strconv.Itoa(value)))
}

func (f fibonacci) getValue(pos int) int {
	if pos == 0 || pos == 1 {
		return 1
	}

	return f.getValue(pos-1) + f.getValue(pos-2)
}

func main() {
	f := fibonacci{}
	router := mux.NewRouter()

	router.Handle("/fib/{value}", f).Methods("GET")
	router.Handle("/fibonacci/{value}", f).Methods("GET")

	http.ListenAndServe(":8000", router)
}
