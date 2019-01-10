package api

import (
	"encoding/json"
	"net/http"
)

type MoveRequest struct {
	Player int   `json:"player"`
	Move   []int `json:"move"`
}

type MoveResponse struct {
	Valid  bool    `json:"valid"`
	Winner int     `json:"winner"`
	Board  [][]int `json:"board"`
}

type IdRequest struct {
	ID string `json:"id"`
}

type IdResponse struct {
	PlayerNum int `json:"playernum"`
}

func NewMoveRequest(req *http.Request) (*MoveRequest, error) {
	d := MoveRequest{}
	err := json.NewDecoder(req.Body).Decode(&d)
	return &d, err
}

func NewIdRequest(req *http.Request) (*IdRequest, error) {
	d := IdRequest{}
	err := json.NewDecoder(req.Body).Decode(&d)
	return &d, err
}
