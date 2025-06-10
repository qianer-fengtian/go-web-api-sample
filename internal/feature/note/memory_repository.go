package note

import (
	"context"
	"errors"
	"sync"

	"example.com/go-web-api-sample/ogen"
)

var (
	ErrNoteNotFound = errors.New("note not found")
)

type MemoryRepository struct {
	notes map[int64]ogen.ModelsNote
	id    int64
	mux   sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		notes: make(map[int64]ogen.ModelsNote),
		id:    1,
	}
}

func (r *MemoryRepository) Create(ctx context.Context, note *ogen.ModelsNote) (*ogen.ModelsNote, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	createdNote := *note
	createdNote.ID = r.id
	r.notes[r.id] = createdNote
	r.id++

	return &createdNote, nil
}

func (r *MemoryRepository) GetByID(ctx context.Context, id int64) (*ogen.ModelsNote, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	note, exists := r.notes[id]
	if !exists {
		return nil, ErrNoteNotFound
	}

	return &note, nil
}

func (r *MemoryRepository) List(ctx context.Context) ([]ogen.ModelsNote, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	notes := make([]ogen.ModelsNote, 0, len(r.notes))
	for _, note := range r.notes {
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *MemoryRepository) Delete(ctx context.Context, id int64) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	if _, exists := r.notes[id]; !exists {
		return ErrNoteNotFound
	}

	delete(r.notes, id)
	return nil
}