package storage

import (
	"database/sql"
	"fmt"
	"log"
	
	"github.com/VadimMLG/MicroService/internal/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitPostgres(connStr string) error {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	
	return db.Ping()
}

func CreatePost(post models.Post) error {
	query := `INSERT INTO posts 
		(id_profile, photo_link, likes_count, comments_count, time, link) 
		VALUES ($1, $2, $3, $4, $5, $6)`
	
	_, err := db.Exec(query,
		post.ProfileID,
		post.PhotoLink,
		post.LikesCount,
		post.CommentsCount,
		post.Time,
		post.Link,
	)
	
	return err
}

func GetUserFeed(userID int) ([]models.Post, error) {
	query := `
	SELECT p.id, p.id_profile, p.photo_link, p.likes_count, 
	       p.comments_count, p.time, p.link 
	FROM posts p
	WHERE p.id_profile IN (
		SELECT id_friend FROM friends WHERE id_profile = $1
		UNION
		SELECT id_follower FROM followers WHERE id_profile = $1
	)
	ORDER BY p.time DESC
	LIMIT 20`
	
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var posts []models.Post
	for rows.Next() {
		var p models.Post
		err := rows.Scan(
			&p.ID,
			&p.ProfileID,
			&p.PhotoLink,
			&p.LikesCount,
			&p.CommentsCount,
			&p.Time,
			&p.Link,
		)
		if err != nil {
			log.Println("Error scanning post:", err)
			continue
		}
		posts = append(posts, p)
	}
	
	return posts, nil
}
