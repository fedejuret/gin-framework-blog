package post

import (
	"time"
)

type Post struct {
	Id      *int      `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Image   *string   `json:"image"`
	Status  uint8     `json:"status"`
	Created time.Time `json:"created_at"`
}
