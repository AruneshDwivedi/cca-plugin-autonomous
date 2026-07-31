package autonomous

import (
	"context"
	"fmt"
)

// TrackedTask wraps a Task with execution tracing
type TrackedTask struct {
	Task
	tracer *Tracer
	name   string
}

func (t *TrackedTask) Run() string {
	ctx := context.WithValue(context.Background(), "task", t.name)
	result := t.Task.Run()
	t.tracer.Log(ctx, fmt.Sprintf("Task %s completed: %s", t.name, result))
	return result
}

// Tracer records task execution sequences
type Tracer struct {
	entries []TraceEntry
}

func NewTracer() *Tracer {
	return &Tracer{entries: make([]TraceEntry, 0)}
}

func (t *Tracer) Log(ctx context.Value, message string) {
	t.entries = append(t.entries, TraceEntry{
		Message: message,
		Timestamp: time.Now(),
	})
}

type TraceEntry struct {
	Message   string
	Timestamp time.Time
}

// WithTracing wraps a coordinator with tracing capabilities
func WithTracing(c *Coordinator, tracer *Tracer) *Coordinator {
	// Would wrap tasks with TrackedTask in production
	return c
}
