package domain

import (
	"errors"
	"time"
)

type Board struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type NewBoardParams struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
}

func NewBoard(params *NewBoardParams) (*Board, error) {
	if len(params.Name) <= 0 {
		return nil, errors.New("Invalid empty name")
	}

	return &Board{
		ID:          params.ID,
		Name:        params.Name,
		Description: params.Description,
		CreatedAt:   params.CreatedAt,
	}, nil
}
