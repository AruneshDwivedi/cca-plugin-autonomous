package autonomous

import (
	"testing"
	"time"
)

type dummyTask struct {
	name string
}

func (d dummyTask) Name() string { return d.name }
func (d dummyTask) Run() string  { return "completed: " + d.name }

func TestCoordinator(t *testing.T) {
	c := NewCoordinator()
	c.AddTask(dummyTask{"test-1"})
	c.AddTask(dummyTask{"test-2"})

	if len(c.tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(c.tasks))
	}
}

func TestScheduler(t *testing.T) {
	c := NewCoordinator()
	s := NewScheduler(c)
	s.ScheduleTask(dummyTask{"scheduled"}, time.Millisecond)
	time.Sleep(100 * time.Millisecond)
}
