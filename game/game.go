package game

import (
	"github.com/otonnesen/tictactoe/api"
)

type Game struct {
	Player int
	Board  [][]int
}

func New() *Game {
	return &Game{1, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}}
}

func (g Game) CheckMove(m *api.MoveRequest) bool {
	if m.Player != g.Player {
		return false
	}

	if g.Board[m.Move[0]][m.Move[1]] != 0 {
		return false
	}

	return true
}

func (g *Game) applyMove(m *api.MoveRequest) {
	g.Board[m.Move[0]][m.Move[1]] = g.Player
	if g.Player == 1 {
		g.Player = 2
	} else {
		g.Player = 1
	}

}
