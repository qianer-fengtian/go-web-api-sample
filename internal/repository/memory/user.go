package memory

import (
	"context"
	"sync"
	"time"

	"example.com/go-web-api-sample/ogen"
	"example.com/go-web-api-sample/internal/repository"
)

type UserRepository struct {
	users map[int64]ogen.ModelsUser
	id    int64
	mux   sync.RWMutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[int64]ogen.ModelsUser),
		id:    1,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *ogen.ModelsUser) (*ogen.ModelsUser, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	createdUser := *user
	createdUser.ID = r.id
	
	// Set createdAt if not provided
	if !createdUser.CreatedAt.IsSet() {
		createdUser.CreatedAt = ogen.NewOptDateTime(time.Now())
	}
	
	r.users[r.id] = createdUser
	r.id++

	return &createdUser, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*ogen.ModelsUser, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, repository.ErrUserNotFound
	}

	return &user, nil
}

func (r *UserRepository) List(ctx context.Context) ([]ogen.ModelsUser, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	users := make([]ogen.ModelsUser, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	if _, exists := r.users[id]; !exists {
		return repository.ErrUserNotFound
	}

	delete(r.users, id)
	return nil
}