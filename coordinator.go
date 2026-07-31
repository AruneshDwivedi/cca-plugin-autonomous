package autonomous

// Coordinator manages execution flow across multiple skills and modules.
type Coordinator struct {
	tasks   []Task
}

func NewCoordinator() *Coordinator {
	return &Coordinator{tasks: make([]Task, 0)}
}

func (c *Coordinator) AddTask(t Task) {
	c.tasks = append(c.tasks, t)
}

func (c *Coordinator) Execute() {
	for _, t := range c.tasks {
		t.Run()
	}
}

// Task represents an autonomous operation
type Task interface {
	Run() string
	Name() string
}
