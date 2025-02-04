package domain

import "time"

type User struct {
	ID        uint
	Name      string
	Email     string
	Active    bool
	CreatedAt time.Time
}
