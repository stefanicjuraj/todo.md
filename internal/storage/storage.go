package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type Storage struct {
	filePath string
}

func NewStorage() *Storage {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}
	
	filePath := filepath.Join(homeDir, ".todo.json")
	return &Storage{filePath: filePath}
}

func (s *Storage) LoadTodos() ([]Todo, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, err
	}

	var todos []Todo
	if len(data) == 0 {
		return []Todo{}, nil
	}

	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}

	return todos, nil
}

func (s *Storage) SaveTodos(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal todos: %w", err)
	}

	err = os.WriteFile(s.filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write todos: %w", err)
	}

	return nil
}

func (s *Storage) renumberTodos(todos []Todo) []Todo {
	for i := range todos {
		todos[i].ID = i + 1
	}
	return todos
}

func (s *Storage) AddTodo(description string) error {
	todos, err := s.LoadTodos()
	if err != nil {
		return err
	}

	newID := len(todos) + 1

	newTodo := Todo{
		ID:          newID,
		Description: description,
		Completed:   false,
	}

	todos = append(todos, newTodo)
	return s.SaveTodos(todos)
}

func (s *Storage) ListTodos() ([]Todo, error) {
	return s.LoadTodos()
}

func (s *Storage) RemoveTodo(id int) error {
	todos, err := s.LoadTodos()
	if err != nil {
		return err
	}

	found := false
	newTodos := make([]Todo, 0, len(todos))
	for _, todo := range todos {
		if todo.ID != id {
			newTodos = append(newTodos, todo)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	newTodos = s.renumberTodos(newTodos)
	return s.SaveTodos(newTodos)
}

func (s *Storage) MarkTodoDone(id int, done bool) error {
	todos, err := s.LoadTodos()
	if err != nil {
		return err
	}

	found := false
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = done
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	return s.SaveTodos(todos)
}
