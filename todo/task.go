package todo

import (
	"strconv"

	"github.com/lithammer/shortuuid/v3"
)

type Task struct {
	info   string
	id     string
	status bool
}

type Tasks []Task

const (
	ID = iota
	INFO
	STATUS
)

func NewTask(info string, status bool) Task {
	return Task{
		id:     shortuuid.New(),
		info:   info,
		status: status,
	}
}

func (task *Task) GetID() string {
	return task.id
}

func (task *Task) GetInfo() string {
	return task.info
}

func (task *Task) SetStatus(status bool) {
	task.status = status
}

func (task *Task) GetStatus() bool {
	return task.status
}

func Task2SliceString(src Task) []string {
	return []string{
		ID:     src.id,
		INFO:   src.info,
		STATUS: strconv.FormatBool(src.status),
	}
}

func sliceString2Task(src []string) (Task, error) {
	srcStatus, err := strconv.ParseBool(src[STATUS])
	if err != nil {
		return Task{}, ErrInvalidSyntax
	}

	return Task{
		id:     src[ID],
		info:   src[INFO],
		status: srcStatus,
	}, nil
}
