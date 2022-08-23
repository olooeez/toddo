package cmd

import (
	"errors"
	"fmt"

	"github.com/lureevar/toddo/todo"
	"github.com/urfave/cli/v2"
)

func Erase(ctx *cli.Context) error {
	toddo := todo.NewTodo(ctx.String("path"))

	if ctx.Args().Len() != 1 {
		return errors.New("you need to pass one TASK ID argument")
	}

	task, err := toddo.GetTask(ctx.Args().First())
	if err != nil {
		return err
	}

	err = toddo.Delete(task)
	if err != nil {
		return err
	}

	if !ctx.Bool("quiet") {
		fmt.Printf("toddo: task with id of \"%v\" has been deleted\n", task.GetID())
	}

	return nil
}
