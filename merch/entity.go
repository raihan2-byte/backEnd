package merch

import "time"

type Merch struct {
	ID int
	Name string
	Price int
	FileName   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
