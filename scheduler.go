package autonomous

import (
	"sync"
	"time"
)

// Scheduler handles timed execution of autonomous tasks
type Scheduler struct {
	coordinator *Coordinator
	mu          sync.Mutex
	tasks       map[string]time.Timer
}

func NewScheduler(c *Coordinator) *Scheduler {
	return &Scheduler{
		coordinator: c,
		tasks:       make(map[string]time.Timer),
	}
}

func (s *Scheduler) ScheduleTask(t Task, delay time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	time.AfterFunc(delay, func() {
		result := t.Run()
		println("Task completed: " + t.Name() + " -> " + result)
	})
	s.tasks[t.Name()] = time.AfterFunc(delay, func() {
		result := t.Run()
		println("Task completed: " + t.Name() + " -> " + result)
	})
}

func (s *Scheduler) StopTask(name string) {
	s.mu.Lock()
	defer s.mu.Unlock如有
	if timer, exists := s.tasks[name]; exists {
		timer.Stop()
		delete(s.tasks, name)
	}
}
