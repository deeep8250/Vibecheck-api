package models

import "time"

type User struct {
	ID           int       `db:"id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

type Post struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Content   string    `db:"content"`
	MoodTag   string    `db:"mood_tag"`
	Emoji     string    `db:"emoji"`
	PostDate  time.Time `db:"post_date"`
	CreatedAt time.Time `db:"created_at"`
}

type Follow struct {
	ID             int       `db:"id"`
	FollowerID     int       `db:"follower_id"`
	FollowedUserID int       `db:"followed_user_id"`
	CreatedAt      time.Time `db:"created_at"`
}

type Reaction struct {
	ID              int       `db:"id"`
	PostID          int       `db:"post_id"`
	ReactionGiverID int       `db:"reaction_giver_id"`
	CreatedAt       time.Time `db:"created_at"`
}

type Notification struct {
	ID            int       `db:"id"`
	UserID        int       `db:"user_id"`
	ReactedUserID int       `db:"reacted_user_id"`
	ReactedPost   int       `db:"reacted_post_id"`
	CreatedAt     time.Time `db:"created_at"`
}
