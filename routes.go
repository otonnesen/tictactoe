package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/otonnesen/tictactoe/api"
	"github.com/otonnesen/tictactoe/game"
)

var games = make(map[string]*game.Game)

func Root(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
}

func Test(w http.ResponseWriter, req *http.Request) {
	_, err := api.NewMoveRequest(req)
	if err != nil {
		Error.Printf("Bad move request: %v\n", err)
	}
	resp := &api.MoveResponse{true, false}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func Id(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Path[len("/id/"):]
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%+v", games[id])
}

func Start(w http.ResponseWriter, req *http.Request) {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 16)
	var id string
	for {
		for i := range b {
			b[i] = "abcdefghijklmnopqrstuvwxyz1234567890"[rand.Intn(36)]
		}
		id = string(b)
		if _, ok := games[id]; !ok {
			break
		}
	}
	games[id] = game.New()
	Info.Printf("Created game %s", id)
	http.Redirect(w, req, "/id/"+id, http.StatusMovedPermanently)
}
