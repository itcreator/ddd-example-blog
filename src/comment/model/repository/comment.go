package repository

import (
	"comment/model/entity"
	"github.com/google/uuid"
	postEntity "post/model/entity"
)

type Comment interface {
	Find(uuid uuid.UUID) (entity.Comment, error)
	Save(comment entity.Comment) error
	List(post postEntity.Post, start, limit int) ([]entity.Comment, error)
	GetTotal(post postEntity.Post) (int, error)
}
