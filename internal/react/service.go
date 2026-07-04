package react

import (
	"context"
)

type ReactService struct {
	repo ReactRepoInterface
}

func NewReactService(r ReactRepoInterface) *ReactService {
	return &ReactService{
		repo: r,
	}
}

func (r *ReactService) ReactPostServ(ctx context.Context, postID, userID int, emoji string) error {
	err := r.repo.ReactPost(ctx, postID, userID, emoji)
	if err != nil {
		return err
	}
	return nil
}
func (r *ReactService) GetReactionsByPostID(ctx context.Context, postID int) (*PostWithReactions, error) {
	reaction, err := r.repo.GetReactionsByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}

	post, err := r.repo.GetPost(ctx, postID)
	if err != nil {
		return nil, err
	}

	var postReactions PostWithReactions
	for _, r := range reaction {
		postReactions.Reactions = append(postReactions.Reactions, r)
	}
	postReactions.OwnerID = post.UserID
	postReactions.PostID = post.ID
	postReactions.Content = post.Content
	postReactions.MoodTag = post.MoodTag
	postReactions.Emoji = post.Emoji
	postReactions.PostDate = post.PostDate
	postReactions.CreatedAt = post.CreatedAt

	return &postReactions, nil

}

func (r *ReactService) RemoveReact(ctx context.Context, postID, userID int) error {
	err := r.repo.DeleteReact(ctx, userID, postID)
	if err != nil {
		return err
	}
	return nil
}
