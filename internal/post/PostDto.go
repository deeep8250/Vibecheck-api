package post

type CreatePost struct {
	Content string `json:"content" binding:"required"`
	MoodTag string `json:"mood_tag" binding:"required"`
	Emoji   string `json:"emoji" binding:"required"`
}
type UpdatePost struct {
	Content string `json:"content" binding:"required"`
	MoodTag string `json:"mood_tag" binding:"required"`
	Emoji   string `json:"emoji" binding:"required"`
}
