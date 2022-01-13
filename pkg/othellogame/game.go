package othellogame

import (
	"github.com/ArminGh02/othello-bot/pkg/othellogame/internal/cell"
	"github.com/ArminGh02/othello-bot/pkg/othellogame/internal/turn"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Game struct {
	users [2]*tgbotapi.User
	board [BOARD_SIZE][BOARD_SIZE]cell.Cell
	turn  turn.Turn
}

func New(user1, user2 *tgbotapi.User) *Game {
	g := &Game{
		users: [2]*tgbotapi.User{user1, user2},
		turn:  turn.Random(),
	}

	mid := len(g.board)/2 - 1
	g.board[mid][mid] = cell.WHITE
	g.board[mid][mid+1] = cell.BLACK
	g.board[mid+1][mid] = cell.BLACK
	g.board[mid+1][mid+1] = cell.WHITE

	return g
}

func (g *Game) IsTurnOf(user *tgbotapi.User) bool {
	return g.users[g.turn.Int()] == user
}
