package board

import (
	"time"

	"github.com/google/uuid"
	"github.com/nagaremono/puck/internal/domain"
	"github.com/nagaremono/puck/pkg/logger"
)

type BoardUseCase interface {
	CreateBoard(params CreateBoardParams) error
}

type CreateBoardParams struct {
	Name        string
	Description string
}

func Name(name string)

type Board struct {
	Logger *logger.Logger
}

func (b *Board) CreateBoard(params CreateBoardParams) error {
	board, err := domain.NewBoard(&domain.NewBoardParams{
		ID:          uuid.New().String(),
		Name:        params.Name,
		Description: params.Description,
		CreatedAt:   time.Now(),
	})
	b.Logger.Info().Msgf("Created Board %w", board)

	if err != nil {
		return err
	}

	return nil
}
