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
	w.Header().Add("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "start.tmpl")))
	err := tmpl.Execute(w, nil)
	if err != nil {
		Error.Printf("Error executing template: %v", err)
	}
}

func Test(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "game.tmpl")))
	err := tmpl.Execute(w, nil)
	if err != nil {
		Error.Printf("Error executing template: %v", err)
	}
}

func Id(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Cache-Control", "no-cache")
	id := req.URL.Path[len("/id/"):]
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "text/html")
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
		w.Header().Add("Content-Type", "application/json")
		if _, ok := games[id]; !ok {
			resp := &api.MoveResponse{false, 0, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}}
			json.NewEncoder(w).Encode(resp)
			return
		}
		g := games[id]
		m, err := api.NewMoveRequest(req)
		if err != nil {
			// Warning.Printf("Bad move request: %v", err)
			resp := &api.MoveResponse{false, g.Winner, g.Board}
			json.NewEncoder(w).Encode(resp)
			return
		}
		valid := g.CheckMove(m)
		winner := g.CheckVictory()
		resp := &api.MoveResponse{valid, winner, g.Board}
		if valid {
			Info.Printf("Valid move: %s\n%s", id, g)
		} else {
			Info.Printf("Invalid move: %s", id)
		}
		json.NewEncoder(w).Encode(resp)
	}
}

func GetId(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(req.Body)
	// Info.Printf("%s", buf.String())
	m, err := api.NewIdRequest(req)
	if err != nil {
		Warning.Printf("Bad ID request: %v", err)
		resp := &api.IdResponse{0}
		json.NewEncoder(w).Encode(resp)
		return
	}
	id := m.ID
	if _, ok := games[id]; !ok {
		Warning.Printf("Bad ID request: %v", err)
	}
	games[id].Connected++
	resp := &api.IdResponse{games[id].Connected}
	json.NewEncoder(w).Encode(resp)
}

func Start(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Cache-Control", "no-cache")
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
