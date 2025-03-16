package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nagaremono/puck/internal/usecase/board"
)

type BoardRoute struct {
	rootRouter http.Handler
	useCase    board.BoardUseCase
}

func NewBoardRoute(root *mux.Router, useCase board.BoardUseCase) (*BoardRoute, error) {
	boardRoute := &BoardRoute{
		rootRouter: root,
		useCase:    useCase,
	}

	s := root.PathPrefix("/boards").Subrouter()

	s.HandleFunc("/", boardRoute.createBoard)

	return boardRoute, nil
}

func (bc *BoardRoute) createBoard(w http.ResponseWriter, r *http.Request) {
	bc.useCase.CreateBoard(board.CreateBoardParams{
		Name:        "Board A",
		Description: "A new board",
	})
}
