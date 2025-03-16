package persistence

import "github.com/nagaremono/puck/internal/domain"

type BoardRepository interface {
	Create(b *domain.Board) error
}

type BoardInMemRepository struct {
	InMemStore[domain.Board]
}

func (br *BoardInMemRepository) Create(b *domain.Board) error {
	br.Set(b.ID, *b)

	return nil
}
