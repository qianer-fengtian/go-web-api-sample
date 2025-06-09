package main

import (
	"log"
	"net/http"

	"example.com/go-web-api-sample/internal/handler"
	"example.com/go-web-api-sample/ogen"
	"example.com/go-web-api-sample/internal/repository/memory"
	"example.com/go-web-api-sample/internal/service"
)

func main() {
	// リポジトリ層の初期化
	noteRepo := memory.NewNoteRepository()

	// サービス層の初期化
	noteService := service.NewNoteService(noteRepo)

	// ハンドラー層の初期化
	noteHandler := handler.NewNoteHandler(noteService)

	// サーバーの初期化
	s, err := ogen.NewServer(noteHandler)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", s); err != nil {
		log.Fatalln(err)
	}
}