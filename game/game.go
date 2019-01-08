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

func (g *Game) CheckVictory() int{

	if(g.Board[0][0]!=0){
		if(g.Board[0][0] == g.Board[0][1] && g.Board[0][2] == g.Board[0][0]) {
			return g.Board[0][0]
		}
		else {
			return 0
		}
	}

	if(g.Board[1][0]!=0){
		if(g.Board[1][0] == g.Board[1][1] && g.Board[1][2] == g.Board[1][0]) {
			return g.Board[1][0]
		} 
		else {
			return 0
		} 
	}

	if(g.Board[2][0]!=0){
		if(g.Board[2][0] == g.Board[2][1] && g.Board[2][2] == g.Board[2][0]) {
			return g.Board[2][0]
		} 
		else {
			return 0
		}
	}

	if(g.Board[0][0]!=0){
		if(g.Board[0][0] == g.Board[1][0] && g.Board[2][0] == g.Board[0][0]) {
			return g.Board[0][0]
		} 
		else {
			return 0
		}
	}
	
	if(g.Board[0][1]!=0){
		if(g.Board[0][1] == g.Board[1][1] && g.Board[2][1] == g.Board[0][1]) {
			return g.Board[0][1]
		}
		else {
			return 0
		}
	}

	if(g.Board[0][2]!=0){
		if(g.Board[0][2] == g.Board[1][2] && g.Board[2][2] == g.Board[0][2]) {
			return g.Board[0][2]
		}
		else {
			return 0
		}
	}

	if(g.Board[0][0]!=0){
		if(g.Board[0][0] == g.Board[1][1] && g.Board[2][2] == g.Board[0][0]) {
			return g.Board[0][0]
		}
		else {
			return 0
		}
	}
	
	if(g.Board[0][2]!=0){
		if(g.Board[0][2] == g.Board[1][1] && g.Board[2][0] == g.Board[0][2]) {
			return g.Board[0][2]
		}
		else {
			return 0
		}
	}
}
