package mock

import (
	"github.com/google/uuid"
	"post/model/entity"
	"post/model/repository"
)

type postRepository struct {
	storage map[uuid.UUID]entity.Post
}

func NewPostRepository() repository.Post {
	return &postRepository{make(map[uuid.UUID]entity.Post)}
}

func (r *postRepository) Save(post entity.Post) error {
	r.storage[post.GetUUID()] = post

	return nil
}

func (r *postRepository) Find(uuid uuid.UUID) (entity.Post, error) {
	return r.storage[uuid], nil
}

func (r *postRepository) Delete(post entity.Post) error {
	id := post.GetUUID()
	if nil != r.storage[id] {
		delete(r.storage, id)
	}

	return nil
}

func (r *postRepository) List(start, limit int) ([]entity.Post, error) {
	var posts []entity.Post
	for _, post := range r.storage {
		if len(posts) >= limit {
			break
		}
		if post.IsEnabled() {
			posts = append(posts, post)
		}
	}

	return posts, nil
}

func (r *postRepository) GetTotal() (int, error) {
	count := 0
	for _, post := range r.storage {
		if post.IsEnabled() {
			count++
		}
	}

	return count, nil
}
