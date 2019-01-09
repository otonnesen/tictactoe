package game

import (
	"fmt"
	"strings"

	"github.com/otonnesen/tictactoe/api"
)

type Game struct {
	Player int
	Board  [][]int
}

func New() *Game {
	return &Game{1, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}}
}

func (g Game) String() string {
	var b strings.Builder
	for x := range g.Board {
		for y := range g.Board[0] {
			fmt.Fprintf(&b, "%d ", g.Board[x][y])
		}
		fmt.Fprintf(&b, "\n")
	}

	return b.String()
}

func (g *Game) CheckMove(m *api.MoveRequest) bool {
	if m.Player != g.Player {
		return false
	}

	if g.Board[m.Move[0]][m.Move[1]] != 0 {
		return false
	}

	g.applyMove(m)
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

func (g Game) CheckVictory() int {
	for x := range g.Board {
		if g.Board[x][0] == g.Board[x][1] && g.Board[x][1] == g.Board[x][2] {
			return g.Board[x][0]
		}
	}

	for y := range g.Board {
		if g.Board[0][y] == g.Board[1][y] && g.Board[1][y] == g.Board[2][y] {
			return g.Board[0][y]
		}
	}

	if g.Board[0][0] == g.Board[1][1] && g.Board[1][1] == g.Board[2][2] {
		return g.Board[0][0]
	}

	if g.Board[2][0] == g.Board[1][1] && g.Board[1][1] == g.Board[0][2] {
		return g.Board[0][0]
	}

	return 0
}
