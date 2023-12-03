package home

import "time"

type TagLineHome struct {
	ID        int
	Heading   string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}