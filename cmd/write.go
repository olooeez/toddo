package cmd

import (
	"errors"
	"fmt"

	"github.com/lureevar/toddo/todo"
	"github.com/urfave/cli/v2"
)

func Write(ctx *cli.Context) error {
	toddo := todo.NewTodo(ctx.String("path"))

	if ctx.Args().Len() == 0 {
		return errors.New("you need to pass the TASK argument")
	}

	for _, i := range ctx.Args().Slice() {
		task := todo.NewTask(i, false)

		err := toddo.Create(task)
		if err != nil {
			return err
		}

		if !ctx.Bool("quiet") {
			fmt.Printf("toddo: \"%v\" as successfully written\n", i)
		}
	}

	return nil
}
