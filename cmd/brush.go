package cmd

import (
	"errors"
	"fmt"

	"github.com/lureevar/toddo/todo"
	"github.com/urfave/cli/v2"
)

func Brush(ctx *cli.Context) error {
	toddo := todo.NewTodo(ctx.String("path"))

	if ctx.Args().Len() != 0 {
		return errors.New("you can't pass arguments to this command")
	}

	tasks, err := toddo.Read()
	if err != nil {
		return err
	}

	hasChanged := false

	for _, i := range tasks {
		if i.GetStatus() {
			err := toddo.Delete(i)
			if err != nil {
				return err
			}

			hasChanged = true

			if !ctx.Bool("quiet") {
				fmt.Printf("toddo: task with id of \"%v\" has been clear\n", i.GetID())
			}
		}
	}

	if !ctx.Bool("quiet") && hasChanged {
		fmt.Println("toddo: todo list has been cleared")
	} else if !ctx.Bool("quiet") && !hasChanged {
		fmt.Println("toddo: there'r no task done in your todo list")
	}

	return nil
}
