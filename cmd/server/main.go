package main

import (
	"log"
	"net/http"

	"example.com/go-web-api-sample/ogen"
	"example.com/go-web-api-sample/internal/feature/note"
	"example.com/go-web-api-sample/internal/feature/user"
)

func main() {
	// リポジトリ層の初期化
	noteRepo := note.NewMemoryRepository()
	userRepo := user.NewMemoryRepository()

	// サービス層の初期化
	noteService := note.NewService(noteRepo)
	userService := user.NewService(userRepo)

	// ハンドラー層の初期化
	noteHandler := note.NewHandler(noteService)
	userHandler := user.NewHandler(userService)

	// 統合ハンドラーの作成
	combinedHandler := &CombinedHandler{
		noteHandler: noteHandler,
		userHandler: userHandler,
	}

	// サーバーの初期化
	s, err := ogen.NewServer(combinedHandler)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Starting server on :8080")
	log.Println("Available endpoints:")
	log.Println("  Notes: GET/POST /notes, GET/DELETE /notes/{id}")
	log.Println("  Users: GET/POST /users, GET/DELETE /users/{id}")
	if err := http.ListenAndServe(":8080", s); err != nil {
		log.Fatalln(err)
	}
}