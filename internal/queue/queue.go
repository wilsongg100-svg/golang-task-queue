package queue

import (
	"errors"

	taskDomain "github.com/wilsongg100-svg/golang-task-queue/internal/task"
)

type Queue struct {
	tasks chan taskDomain.Task
}

func (q *Queue) Add(task taskDomain.Task) error {
	if task.Status != taskDomain.STATUS_PENDING {
		return errors.New("invalid status")
	}
	q.tasks <- task
	return nil
}

func (q *Queue) Get() (taskDomain.Task, bool) {
	task, ok := <-q.tasks
	return task, ok
}

func (q *Queue) Close() {
	close(q.tasks)
}

func (q *Queue) GetSize() int {
	return len(q.tasks)
}
