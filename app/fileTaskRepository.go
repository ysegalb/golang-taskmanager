package app

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type FileTaskRepository struct {
	filename string
	mutex    sync.RWMutex
}

func NewFileTaskRepository(filename string) (TaskRepository, error) {
	return &FileTaskRepository{
		filename: filename,
	}, nil
}

func (r *FileTaskRepository) readTasks() ([]Task, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	data, err := os.ReadFile(r.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []Task{}, nil
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *FileTaskRepository) writeTasks(tasks []Task) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filename, data, 0644)
}

func (r *FileTaskRepository) Save(task Task) (*Task, error) {
	tasks, err := r.readTasks()
	if err != nil {
		return nil, err
	}

	tasks = append(tasks, task)
	if err := r.writeTasks(tasks); err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *FileTaskRepository) FindAll() ([]Task, error) {
	return r.readTasks()
}

func (r *FileTaskRepository) Update(id int, updatedTask Task) (*Task, error) {
	tasks, err := r.readTasks()
	if err != nil {
		return nil, err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updatedTask
			if err := r.writeTasks(tasks); err != nil {
				return nil, err
			}
			return &updatedTask, nil
		}
	}

	return nil, fmt.Errorf("task with id %d not found", id)
}

func (r *FileTaskRepository) Delete(id int) error {
	tasks, err := r.readTasks()
	if err != nil {
		return err
	}

	var newTasks []Task
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}

	return r.writeTasks(newTasks)
}

func (r *FileTaskRepository) FindById(id int) (*Task, error) {
	tasks, err := r.readTasks()
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}

	return nil, fmt.Errorf("task with id %d not found", id)
}

func (r *FileTaskRepository) FindByDescription(description string) ([]Task, error) {
	tasks, err := r.readTasks()
	if err != nil {
		return nil, err
	}

	var result []Task
	for _, task := range tasks {
		if task.Description == description {
			result = append(result, task)
		}
	}

	return result, nil
}

func (r *FileTaskRepository) FindByStatus(status Status) ([]Task, error) {
	tasks, err := r.readTasks()
	if err != nil {
		return nil, err
	}

	var result []Task
	for _, task := range tasks {
		if task.Status == status {
			result = append(result, task)
		}
	}

	return result, nil
}

func (r *FileTaskRepository) FindByStatusAndDescription(status Status, description string) ([]Task, error) {
	tasks, err := r.readTasks()
	if err != nil {
		return nil, err
	}

	var result []Task
	for _, task := range tasks {
		if task.Status == status && task.Description == description {
			result = append(result, task)
		}
	}

	return result, nil
}

func (r *FileTaskRepository) MarkDone(id int) error {
	task, err := r.FindById(id)
	if err != nil {
		return err
	}

	task.Status = DONE
	_, err = r.Update(id, *task)
	return err
}

func (r *FileTaskRepository) MarkInProgress(id int) error {
	task, err := r.FindById(id)
	if err != nil {
		return err
	}

	task.Status = IN_PROGRESS
	_, err = r.Update(id, *task)
	return err
}

func (r *FileTaskRepository) GetNextID() int {
	tasks, err := r.readTasks()
	if err != nil {
		return 1
	}

	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID + 1
}
