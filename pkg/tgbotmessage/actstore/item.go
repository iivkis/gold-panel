package actstore

import "time"

type Item struct {
	Action    string
	ExpiresIn time.Time
}
