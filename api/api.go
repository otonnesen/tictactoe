package api

import (
	"encoding/json"
	"net/http"
)

type MoveRequest struct {
	ID     string `json:"id"`
	Player int    `json:"player"`
	Move   []int  `json:"move"`
}

type MoveResponse struct {
	Valid bool    `json:"valid"`
	Board [][]int `json:"board"`
	Turn  int     `json:"turn"`
	Done  bool    `json:"done"`
}

func NewMoveRequest(req *http.Request) (*MoveRequest, error) {
	d := MoveRequest{}
	err := json.NewDecoder(req.Body).Decode(&d)
	return &d, err
}
