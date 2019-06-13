package mock

import (
	"comment/model/entity"
	"comment/model/repository"
	"github.com/google/uuid"
	postEntity "post/model/entity"
)

type commentRepository struct {
	storage map[uuid.UUID]entity.Comment
}

func NewCommentRepository() repository.Comment {
	return &commentRepository{make(map[uuid.UUID]entity.Comment)}
}

func (r *commentRepository) Save(comment entity.Comment) error {
	r.storage[comment.GetUUID()] = comment

	return nil
}

func (r *commentRepository) Find(uuid uuid.UUID) (entity.Comment, error) {
	return r.storage[uuid], nil
}

func (r *commentRepository) Delete(comment entity.Comment) error {
	id := comment.GetUUID()
	if nil != r.storage[id] {
		delete(r.storage, id)
	}

	return nil
}

func (r *commentRepository) List(post postEntity.Post, start, limit int) ([]entity.Comment, error) {
	var comments []entity.Comment
	for _, comment := range r.storage {
		if len(comments) >= limit {
			break
		}
		if comment.GetPost().GetUUID() == post.GetUUID() {
			comments = append(comments, comment)
		}
	}

	return comments, nil
}

func (r *commentRepository) GetTotal(post postEntity.Post) (int, error) {
	count := 0
	for _, comment := range r.storage {
		if comment.GetPost().GetUUID() == post.GetUUID() {
			count++
		}
	}

	return count, nil
}
