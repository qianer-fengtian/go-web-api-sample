package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"

	"example.com/go-web-api-sample/ogen"
)

type noteService struct {
	notes map[int64]ogen.ModelsNote
	id    int64
	mux   sync.Mutex
}

func (n *noteService) CreateNote(_ context.Context, req *ogen.ModelsNote) (*ogen.ModelsNote, error) {
	n.mux.Lock()
	defer n.mux.Unlock()

	n.notes[n.id] = *req
	n.id++
	return req, nil
}

func (n *noteService) GetNote(ctx context.Context, params ogen.GetNoteParams) (*ogen.ModelsNote, error) {
	n.mux.Lock()
	defer n.mux.Unlock()

	note, ok := n.notes[params.ID]	
	if !ok {
		return nil, errors.New("メモが見つかりません")
	}

	return &note, nil
}

func main() {
	service := &noteService{
		notes: map[int64]ogen.ModelsNote{},
	}
	s, err := ogen.NewServer(service)
	if err != nil {
		log.Fatalln(err)
	}
	if err := http.ListenAndServe(":8080", s); err != nil {
		log.Fatalln(err)
	}
}