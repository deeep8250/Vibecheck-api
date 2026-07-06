package react

import (
	"time"
)

type ReactionDetail struct {
	ReactionID      int       `db:"id" json:"id"`
	ReactorUsername string    `db:"username" json:"username"`
	Emoji           string    `db:"reaction_emoji" json:"emoji"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}
type PostWithReactions struct {
	PostID    int              `json:"post_id"`
	OwnerID   int              `json:"owner_id"`
	Content   string           `json:"content"`
	MoodTag   string           `json:"mood_tag"`
	Emoji     string           `json:"emoji"`
	PostDate  time.Time        `json:"post_date"`
	CreatedAt time.Time        `json:"created_at"`
	Reactions []ReactionDetail `json:"reactions"`
}
