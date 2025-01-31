package todo

import (
	"sync"

	"github.com/ilyas-zal/todo-go/internal/models"
)

type TodoService struct {
	todos []models.Todo
	mu    sync.Mutex
}

func NewTodoService() *TodoService {
	return &TodoService{
		todos: make([]models.Todo, 0),
	}
}

func (s *TodoService) AddTask(task string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.todos = append(s.todos, models.Todo{Task: task})
}

func (s *TodoService) CompleteTask(index int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if index >= 0 && index < len(s.todos) {
		s.todos[index].Complete = !s.todos[index].Complete
	}
}

func (s *TodoService) GetTasks() []models.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.todos
}
