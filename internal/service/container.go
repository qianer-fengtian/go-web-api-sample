package service

// ServiceContainer holds all application services
// This makes it easy to add new services without changing function signatures
type ServiceContainer struct {
	NoteService *NoteService
	UserService *UserService
	// 新しいサービスはここに追加するだけ
	// ProductService *ProductService
	// OrderService   *OrderService
}

// NewServiceContainer creates a new service container with all services
func NewServiceContainer(
	noteService *NoteService,
	userService *UserService,
) *ServiceContainer {
	return &ServiceContainer{
		NoteService: noteService,
		UserService: userService,
	}
}