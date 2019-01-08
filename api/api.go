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
	Valid  bool `json:"valid"`
	Victor int  `json:"victor"`
}

func NewMoveRequest(req *http.Request) (*MoveRequest, error) {
	d := MoveRequest{}
	err := json.NewDecoder(req.Body).Decode(&d)
	return &d, err
}
