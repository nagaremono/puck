package domain

import "time"

type Ticket struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Items       []TicketItem `json:"items"`
	CreatedAt   time.Time    `json:"createdAt"`
}
