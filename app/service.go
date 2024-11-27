package app

import (
	"fmt"
	"time"
)

type TaskService interface {
	AddTask(description string) (*Task, error)
	ListTasks() ([]Task, error)
	ListTasksByStatus(status Status) ([]Task, error)
	UpdateTask(id int, description string) (*Task, error)
	DeleteTask(id int) error
	MarkTaskInProgress(id int) error
	MarkTaskDone(id int) error
}

type taskService struct {
	taskRepository TaskRepository
}

func NewTaskService(taskRepository TaskRepository) TaskService {
	return &taskService{
		taskRepository: taskRepository,
	}
}

func (s *taskService) AddTask(description string) (*Task, error) {
	task := Task{
		ID:          s.taskRepository.GetNextID(),
		Description: description,
		Status:      TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	savedTask, err := s.taskRepository.Save(task)
	if err != nil {
		return nil, fmt.Errorf("failed to add task: %w", err)
	}
	return savedTask, nil
}

func (s *taskService) ListTasks() ([]Task, error) {
	tasks, err := s.taskRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}
	return tasks, nil
}

func (s *taskService) ListTasksByStatus(status Status) ([]Task, error) {
	tasks, err := s.taskRepository.FindByStatus(status)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks by status: %w", err)
	}
	return tasks, nil
}

func (s *taskService) UpdateTask(id int, description string) (*Task, error) {
	task, err := s.taskRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("task not found: %w", err)
	}
	task.Description = description
	task.UpdatedAt = time.Now()
	
	updatedTask, err := s.taskRepository.Update(id, *task)
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}
	return updatedTask, nil
}

func (s *taskService) DeleteTask(id int) error {
	if err := s.taskRepository.Delete(id); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

func (s *taskService) MarkTaskInProgress(id int) error {
	if err := s.taskRepository.MarkInProgress(id); err != nil {
		return fmt.Errorf("failed to mark task as in progress: %w", err)
	}
	return nil
}

func (s *taskService) MarkTaskDone(id int) error {
	if err := s.taskRepository.MarkDone(id); err != nil {
		return fmt.Errorf("failed to mark task as done: %w", err)
	}
	return nil
}
