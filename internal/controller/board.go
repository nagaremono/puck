package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nagaremono/puck/internal/usecase/board"
	"github.com/nagaremono/puck/pkg/logger"
)

type BoardRoute struct {
	rootRouter http.Handler
	useCase    board.BoardUseCase

	logger *logger.Logger
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
	var body PostBoardsBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		bc.logger.Logger.Error().Err(err).Msg("Error parsing body")
	}

	bc.useCase.CreateBoard(board.CreateBoardParams{
		Name:        body.Name,
		Description: body.Description,
	})
}
