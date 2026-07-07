package worker

import (
	"context"
	"log"

	"github.com/deeep8250/vibecheck-api/internal/react"
)

type ReactJob struct {
	PostID int
	UserID int
	Emoji  string
}

type ReactWorker struct {
	channel chan ReactJob
	repo    *react.ReactRepository
}

func NewReactWorker(Repo *react.ReactRepository) *ReactWorker {
	return &ReactWorker{
		channel: make(chan ReactJob, 100),
		repo:    Repo,
	}
}

func (w *ReactWorker) Start() {
	go func() {
		for job := range w.channel {
			Ctx := context.Background()

			err := w.repo.ReactPost(Ctx, job.PostID, job.UserID, job.Emoji)
			if err != nil {
				log.Println("upvote worker error :", err.Error())
			}
		}
	}()
}

func (w *ReactWorker) Submit(postID, userID int, emoji string) {
	w.channel <- ReactJob{

		PostID: postID,
		UserID: userID,
		Emoji:  emoji,
	}
}
