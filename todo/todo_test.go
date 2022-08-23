package todo

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

const (
	SampleIn              = "./samples/in.csv"
	SampleOut             = "./samples/out.csv"
	SampleOutBroken       = "./samples/out_broken.csv"
	SampleOutBrokenSyntax = "./samples/out_broken_syntax.csv"
	StupidPath            = "./path/that/doesnt/exist"
)

func TestNewTodo(t *testing.T) {
	want := &Todo{
		path: SampleIn,
	}

	got := NewTodo(SampleIn)

	if got.path != want.path {
		t.Errorf("NewTodo() => got.path(%v) != want.path(%v)", got.path, want.path)
	}
}

func TestCreate(t *testing.T) {
	task := NewTask("Buy Milk", false)
	todo := NewTodo(SampleIn)

	_ = todo.Create(task)

	got, _ := ioutil.ReadFile(SampleIn)
	want := fmt.Sprintf("%s,%s,%v\n", task.id, task.info, task.status)

	if string(got) != want {
		t.Errorf("(*Todo).Create(%v) => got (%v), but want (%v)", task, string(got), want)
	}

	_ = os.Truncate(SampleIn, 0)
}

func TestCreateFileError(t *testing.T) {
	task := NewTask("Buy Milk", false)
	todo := NewTodo(StupidPath)

	err := todo.Create(task)
	if !errors.Is(err, ErrCouldntOpenTodo) {
		t.Errorf("(*Todo).Create(%v) => %v", task, err)
	}
}

func TestRead(t *testing.T) {
	todo := NewTodo(SampleOut)
	tasks, _ := todo.Read()

	file, _ := os.Open(SampleOut)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	tasksCount := 0

	for scanner.Scan() {
		rawTask := strings.Split(scanner.Text(), ",")

		for i, j := range rawTask {
			if i == ID && j != tasks[tasksCount].id {
				t.Errorf("(*Todo).Read() => j[ID](%v) != tasks[%d].id(%v)", j, tasksCount, tasks[tasksCount].id)
			}

			if i == INFO && j != tasks[tasksCount].info {
				t.Errorf("(*Todo).Read() => j[INFO](%v) != tasks[%d].info(%v)", j, tasksCount, tasks[tasksCount].info)
			}

			if i == STATUS && j != strconv.FormatBool(tasks[tasksCount].status) {
				t.Errorf("(*Todo).Read() => j[STATUS](%v) != tasks[%d].status(%v)", j, tasksCount, tasks[tasksCount].status)
			}
		}

		tasksCount++
	}
}

func TestReadFileError(t *testing.T) {
	todo := NewTodo(StupidPath)

	_, err := todo.Read()
	if !errors.Is(err, ErrCouldntOpenTodo) {
		t.Errorf("(*Todo).Read() => %v", err)
	}
}

func TestReadParseError(t *testing.T) {
	todo := NewTodo(SampleOutBroken)

	_, err := todo.Read()
	if !errors.Is(err, ErrCouldntParseTodo) {
		t.Errorf("(*Todo).Read() => %v", err)
	}
}

func TestReadParseBoolError(t *testing.T) {
	todo := NewTodo(SampleOutBrokenSyntax)

	_, err := todo.Read()
	if !errors.Is(err, ErrInvalidSyntax) {
		t.Errorf("(*Todo).Read() => %v", err)
	}
}

func TestDelete(t *testing.T) {
	todo := NewTodo(SampleIn)

	tasks := Tasks{
		NewTask("Buy some milk", false),
		NewTask("Make tests for toddo", true),
		NewTask("Make some tea", false),
	}

	for _, i := range tasks {
		_ = todo.Create(i)
	}

	oldTasks, _ := todo.Read()
	_ = todo.Delete(tasks[0])
	newTasks, _ := todo.Read()

	if reflect.DeepEqual(oldTasks, newTasks) {
		t.Errorf("(*Todo).Delete(%v) => oldTasks(%v) == newTasks(%v)", tasks[0], oldTasks, newTasks)
	}

	_ = os.Truncate(SampleIn, 0)
}

func TestDeleteFileError(t *testing.T) {
	todo := NewTodo(StupidPath)
	task := NewTask("Buy some milk", false)

	err := todo.Delete(task)
	if !errors.Is(err, ErrCouldntOpenTodo) {
		t.Errorf("(*Todo).Delete(%v) => %v", task, err)
	}
}

func TestDeleteGetTasks(t *testing.T) {
	todo := NewTodo(SampleOutBroken)
	task := NewTask("Buy some milk", false)

	err := todo.Delete(task)
	if !errors.Is(err, ErrCouldntParseTodo) {
		t.Errorf("(*Todo).Delete(%v) => %v", task, err)
	}
}

func TestDeleteCantFindTask(t *testing.T) {
	todo := NewTodo(SampleIn)
	tasks := Tasks{
		NewTask("Buy some milk", false),
		NewTask("Make tests for toddo", true),
		NewTask("Make some tea", false),
	}

	for _, i := range tasks {
		_ = todo.Create(i)
	}

	nonExistingTask := NewTask("This task doesn't exist in the todo file", false)

	err := todo.Delete(nonExistingTask)
	if !errors.Is(err, ErrCantFindTask) {
		t.Errorf("(*Todo).Delete(%v) => %v", nonExistingTask, err)
	}

	_ = os.Truncate(SampleIn, 0)
}

func TestUpdate(t *testing.T) {
	todo := NewTodo(SampleIn)

	tasks := Tasks{
		NewTask("Buy some milk", false),
		NewTask("Make tests for toddo", true),
		NewTask("Make some tea", false),
	}

	for _, i := range tasks {
		_ = todo.Create(i)
	}

	newTask := tasks[0]
	newTask.status = true

	oldTasks, _ := todo.Read()
	_ = todo.Update(tasks[0], newTask)
	newTasks, _ := todo.Read()

	if reflect.DeepEqual(oldTasks, newTasks) {
		t.Errorf("(*Todo).Update(%v, %v) => oldTasks(%v) == newTasks(%v)", tasks[0], newTask, oldTasks, newTasks)
	}

	_ = os.Truncate(SampleIn, 0)
}

func TestUpdateFileError(t *testing.T) {
	todo := NewTodo(StupidPath)

	oldTask := NewTask("Buy some milk", false)
	newTask := oldTask
	newTask.status = true

	err := todo.Update(oldTask, newTask)
	if !errors.Is(err, ErrCouldntOpenTodo) {
		t.Errorf("(*Todo).Delete(%v, %v) => %v", oldTask, newTask, err)
	}
}

func TestUpdateGetTasks(t *testing.T) {
	todo := NewTodo(SampleOutBroken)

	oldTask := NewTask("Buy some milk", false)
	newTask := oldTask
	newTask.status = true

	err := todo.Update(oldTask, newTask)
	if !errors.Is(err, ErrCouldntParseTodo) {
		t.Errorf("(*Todo).Delete(%v, %v) => %v", oldTask, newTask, err)
	}
}

func TestUpdateCantFindTask(t *testing.T) {
	todo := NewTodo(SampleIn)

	tasks := Tasks{
		NewTask("Buy some milk", false),
		NewTask("Make tests for toddo", true),
		NewTask("Make some tea", false),
	}

	for _, i := range tasks {
		_ = todo.Create(i)
	}

	nonExistingTask := NewTask("This task doesn't exist in the todo file", false)

	err := todo.Update(tasks[0], nonExistingTask)
	if !errors.Is(err, ErrCantFindTask) {
		t.Errorf("(*Todo).Delete(%v, %v) => %v", tasks[0], nonExistingTask, err)
	}

	_ = os.Truncate(SampleIn, 0)
}
