package home

import "time"

type TagLineHome struct {
	ID        int
	Heading   string
	Text      string
	CreatedAt time.Time
	UpdatesAt time.Time
}