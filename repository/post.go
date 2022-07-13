package repository

import (
	"context"
	"github.com/harrycoa/apirest-websockets.git/models"
)

type PostRepository interface {
	InsertPost(ctx context.Context, post *models.Post) error
}

var implementationPost PostRepository

func SetPostRepository(repository PostRepository) {
	implementationPost = repository
}

func ClosePost() error {
	return implementation.Close()
}

func InsertPost(ctx context.Context, post *models.Post) error {
	return implementationPost.InsertPost(ctx, post)
}
