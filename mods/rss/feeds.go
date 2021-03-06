package rss

import (
	"time"
)

// Feeds are the allowed and available web feeds that users can subscribe
// to.
type Feed struct {
	Id            uint64
	Name          string
	Url           string
	LastPublished time.Time
}
