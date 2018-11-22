package repository

import (
	"github.com/google/uuid"
	"post/model/entity"
)

type Post interface {
	Find(uuid uuid.UUID) (entity.Post, error)
	Save(post entity.Post) error
}
