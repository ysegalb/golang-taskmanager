package app

type TaskRepository interface {
	Save(task Task) (*Task, error)
	FindAll() ([]Task, error)
	Update(id int, task Task) (*Task, error)
	Delete(id int) error
	MarkDone(id int) error
	MarkInProgress(id int) error
	FindById(id int) (*Task, error)
	FindByDescription(description string) ([]Task, error)
	FindByStatus(status Status) ([]Task, error)
	GetNextID() int
}
