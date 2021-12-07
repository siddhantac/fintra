package domain

import "time"

type Account struct {
	ID              string
	Name            string
	Balance         int64
	StartingBalance int64
	Created         time.Time
	Updated         time.Time
}
