package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/otonnesen/tictactoe/api"
)

func main() {
	http.HandleFunc("/", Root)
	http.HandleFunc("/game", Game)

	http.ListenAndServe(":8080", nil)
}

func Root(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello World")
}

func Game(w http.ResponseWriter, req *http.Request) {
	_, err := api.NewMoveRequest(req)
	if err != nil {
		panic(err)
	}
	resp := &api.MoveResponse{true, [][]int{[]int{0}}, 1, false}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
