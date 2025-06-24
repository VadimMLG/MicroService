package models

import "time"

type Post struct {
	ID            uint      `json:"id"`
	ProfileID     uint      `json:"profile_id"`
	PhotoLink     string    `json:"photo_link"`
	LikesCount    int       `json:"likes_count"`
	CommentsCount int       `json:"comments_count"`
	Time          time.Time `json:"time"`
	Link          string    `json:"link"`
}
