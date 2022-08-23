package todo

import (
	"encoding/csv"
	"errors"
	"os"
)

type Todo struct {
	path string
}

var (
	ErrCouldntOpenTodo  = errors.New("couldn't open the toddo todo list file")
	ErrCouldntParseTodo = errors.New("couldn't parse the toddo todo list file and it's tasks")
	ErrInvalidSyntax    = errors.New("couldn't parse the toddo todo list file because of invalid syntax")
	ErrCantFindTask     = errors.New("can't find the task (task doesn't exist)")
	ErrTaskEqual        = errors.New("task selected is equal to the old one")
)

func NewTodo(path string) *Todo {
	return &Todo{
		path: path,
	}
}

func (todo *Todo) Create(task Task) error {
	file, err := os.OpenFile(todo.path, os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return ErrCouldntOpenTodo
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	taskRaw := Task2SliceString(task)

	_ = writer.Write(taskRaw)
	writer.Flush()

	return nil
}

func (todo *Todo) Read() (Tasks, error) {
	file, err := os.OpenFile(todo.path, os.O_RDONLY, 0o644)
	if err != nil {
		return nil, ErrCouldntOpenTodo
	}

	defer file.Close()

	reader := csv.NewReader(file)

	rawTasks, err := reader.ReadAll()
	if err != nil {
		return nil, ErrCouldntParseTodo
	}

	tasks := make(Tasks, 0)

	for _, i := range rawTasks {
		task, err := sliceString2Task(i)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (todo *Todo) Delete(task Task) error {
	file, err := os.OpenFile(todo.path, os.O_RDWR, 0o644)
	if err != nil {
		return ErrCouldntOpenTodo
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	tasks, err := todo.Read()
	if err != nil {
		return err
	}

	newTasks := make([][]string, 0)

	for _, i := range tasks {
		if i.id != task.id {
			rawTask := Task2SliceString(i)
			newTasks = append(newTasks, rawTask)
		}
	}

	if len(newTasks) == len(tasks) {
		return ErrCantFindTask
	}

	_ = os.Truncate(todo.path, 0)
	_ = writer.WriteAll(newTasks)

	return nil
}

func (todo *Todo) Update(oldTask, newTask Task) error {
	file, err := os.OpenFile(todo.path, os.O_RDWR, 0o644)
	if err != nil {
		return ErrCouldntOpenTodo
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	oldTasks, err := todo.Read()
	if err != nil {
		return err
	}

	newTasks := make([][]string, 0)
	tasksEquals := true

	for _, i := range oldTasks {
		if i.id == newTask.id && i == newTask {
			return ErrTaskEqual
		}

		if i.id == newTask.id {
			rawNewTask := Task2SliceString(newTask)
			newTasks = append(newTasks, rawNewTask)
			tasksEquals = false
		} else {
			rawTask := Task2SliceString(i)
			newTasks = append(newTasks, rawTask)
		}
	}

	if tasksEquals {
		return ErrCantFindTask
	}

	_ = os.Truncate(todo.path, 0)
	_ = writer.WriteAll(newTasks)

	return nil
}

func (todo *Todo) GetTask(id string) (Task, error) {
	tasks, err := todo.Read()
	if err != nil {
		return Task{}, err
	}

	for _, i := range tasks {
		if i.id == id {
			return i, nil
		}
	}

	return Task{}, ErrCantFindTask
}
