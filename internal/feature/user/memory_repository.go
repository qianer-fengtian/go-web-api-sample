package user

import (
	"context"
	"errors"
	"sync"
	"time"

	"example.com/go-web-api-sample/ogen"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type MemoryRepository struct {
	users map[int64]ogen.ModelsUser
	id    int64
	mux   sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		users: make(map[int64]ogen.ModelsUser),
		id:    1,
	}
}

func (r *MemoryRepository) Create(ctx context.Context, user *ogen.ModelsUser) (*ogen.ModelsUser, error) {
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

func (r *MemoryRepository) GetByID(ctx context.Context, id int64) (*ogen.ModelsUser, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}

	return &user, nil
}

func (r *MemoryRepository) List(ctx context.Context) ([]ogen.ModelsUser, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	users := make([]ogen.ModelsUser, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

func (r *MemoryRepository) Delete(ctx context.Context, id int64) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	if _, exists := r.users[id]; !exists {
		return ErrUserNotFound
	}

	delete(r.users, id)
	return nil
}