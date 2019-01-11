package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"path/filepath"
	"time"

	"github.com/otonnesen/tictactoe/api"
	"github.com/otonnesen/tictactoe/game"
)

var games = make(map[string]*game.Game)

func Root(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
}

func Test(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "game.tmpl")))
	err := tmpl.Execute(w, nil)
	if err != nil {
		Error.Printf("Error executing template: %v", err)
	}
}

func Id(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Path[len("/id/"):]
	switch req.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "text/html")
		if _, ok := games[id]; ok {
			// fmt.Fprintf(w, "%s", g)
			fp := filepath.Join("templates", "game.tmpl")
			tmpl := template.Must(template.ParseFiles(fp))
			err := tmpl.Execute(w, nil)
			if err != nil {
				Error.Printf("Error executing template: %v", err)
			}
		} else {
			fmt.Fprintf(w, "Not a valid game")
		}
	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		if _, ok := games[id]; !ok {
			resp := &api.MoveResponse{false, 0}
			json.NewEncoder(w).Encode(resp)
			return
		}
		m, err := api.NewMoveRequest(req)
		if err != nil {
			Error.Printf("Bad move request: %v\n", err)
			resp := &api.MoveResponse{false, 0}
			json.NewEncoder(w).Encode(resp)
			return
		}
		g := games[id]
		valid := g.CheckMove(m)
		done := g.CheckVictory()
		resp := &api.MoveResponse{valid, done}
		if valid {
			Info.Printf("Valid move: %s\n%s", id, g)
		} else {
			Info.Printf("Invalid move: %s", id)
		}
		json.NewEncoder(w).Encode(resp)
	}
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
	Info.Printf("Created game %s:\n%v", id, games[id])
	http.Redirect(w, req, "/id/"+id, http.StatusMovedPermanently)
}
