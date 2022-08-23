package cmd

import (
	"errors"
	"fmt"

	"github.com/lureevar/toddo/todo"
	"github.com/urfave/cli/v2"
)

func Mark(ctx *cli.Context) error {
	toddo := todo.NewTodo(ctx.Path("path"))

	if ctx.Args().Len() != 1 {
		return errors.New("you need to pass one TASK ID argument")
	}

	oldTask, err := toddo.GetTask(ctx.Args().First())
	if err != nil {
		return err
	}

	newTask := oldTask
	newTask.SetStatus(true)

	err = toddo.Update(oldTask, newTask)
	if err != nil {
		return err
	}

	if !ctx.Bool("quiet") {
		fmt.Printf("toddo: task with id of \"%v\" had its status changed from false to true\n", oldTask.GetID())
	}

	return nil
}
