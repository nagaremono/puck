package controller

import (
	"encoding/json"
	"net/http"
)

type PostBoardsBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func FromBody(r *http.Request) *PostBoardsBody {
	var board PostBoardsBody

	json.NewDecoder(r.Body).Decode(&board)

	return &board
}
