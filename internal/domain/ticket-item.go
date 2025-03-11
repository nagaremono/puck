package domain

import "time"

type TicketItem struct {
	ID        string    `json:"id"`
	BoardID   string    `json:"boardId"`
	Idx       int       `json:"idx"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
