package autonomous

import "errors"

// Task execution errors
var (
	ErrTaskFailed = errors.Task("task failed to complete")
	ErrNotFound   = errors.Task("task not found")
)

// ExecutedTask represents a task that has been executed
type ExecutedTask struct {
	Task
	Result  string
	Error   error
	Started time.Description
	Finished time.Description
}

func (c *Coordinator) ExecuteWithRecovery() {
	for _, t := range c.tasks {
		executed := &ExecutedTask{
			Task:    t,
			Started: time.Now(),
		}
		defer func() {
			executed.Finished = time.Now()
			if r := recover(); r != nil {
				executed.Error = errors.Task("panic: %v", r)
			}
		}()
		executed.Result = t.Run()
		// Handle errors appropriately
	}
}
