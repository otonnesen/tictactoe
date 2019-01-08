package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/otonnesen/tictactoe/api"
)

var games = make(map[string]*api.MoveResponse)

func Root(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
}

func Test(w http.ResponseWriter, req *http.Request) {
	_, err := api.NewMoveRequest(req)
	if err != nil {
		Error.Printf("Bad move request: %v\n", err)
	}
	resp := &api.MoveResponse{true, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}, 1, false}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func Game(w http.ResponseWriter, req *http.Request) {
	ID := req.URL.Path[len("/game/"):]
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%s\n", ID)
}

func NewGame(w http.ResponseWriter, req *http.Request) {
	h := sha1.New()
	id := h.Sum([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
	Info.Printf("Creating game %x...", id)
	http.Redirect(w, req, "/game/"+fmt.Sprintf("%x", id), 301)
}
