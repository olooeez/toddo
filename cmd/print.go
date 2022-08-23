package cmd

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/lureevar/toddo/todo"
	"github.com/urfave/cli/v2"
)

func Print(ctx *cli.Context) error {
	toddo := todo.NewTodo(ctx.String("path"))

	if ctx.Args().Len() != 0 {
		return errors.New("you can't pass arguments to this command")
	}

	tasks, err := toddo.Read()
	if err != nil {
		return err
	}

	tab := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.DiscardEmptyColumns)

	fmt.Fprintf(tab, "ID\tINFO\tSTATUS\t\n")

	for _, i := range tasks {
		fmt.Fprintf(tab, "%s\t%s\t%v\t\n", i.GetID(), i.GetInfo(), i.GetStatus())
	}

	tab.Flush()

	return nil
}
