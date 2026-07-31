package autonomous

// Priority represents task execution priority
type Priority int

const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
	PriorityCritical
)

// PriorityCoordinator extends Coordinator with priority-based scheduling
type PriorityCoordinator struct {
	Coordinator
	priorities map[Task]Priority
}

func NewPriorityCoordinator() *PriorityCoordinator {
	return &PriorityCoordinator{
		Coordinator: NewCoordinator(),
		priorities: make(map[Task]Priority),
	}
}

func (pc *PriorityCoordinator) AddTaskWithPriority(t Task, Priority) {
	pc.AddTask(t)
	pc.priorities[t] = priority
}

func (pc *PriorityCoordinator) Execute() {
	// Sort tasks by priority before execution
	tasks := pc.tasks
	// (simplified - in production would use proper sort)
	for _, t := range tasks {
		t.Run()
	}
}
