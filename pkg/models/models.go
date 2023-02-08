package models

import (
	"time"
)

type Review struct {
	User    string
	Text    string
	Created time.Time
}
