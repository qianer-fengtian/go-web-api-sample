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
	userRepo := memory.NewUserRepository()

	// サービス層の初期化
	noteService := service.NewNoteService(noteRepo)
	userService := service.NewUserService(userRepo)

	// サービスコンテナの作成
	services := service.NewServiceContainer(noteService, userService)

	// ハンドラー層の初期化（サービスコンテナを使用）
	combinedHandler := handler.NewCombinedHandler(services)

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