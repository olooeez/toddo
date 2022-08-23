package main

import (
	"log"
	"os"

	"github.com/lureevar/toddo/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	defaultHome, _ := os.UserHomeDir()
	defaultPathTodo := defaultHome + string(os.PathSeparator) + ".toddo.csv"

	if _, ok := os.Stat(defaultPathTodo); os.IsNotExist(ok) {
		_, _ = os.Create(defaultPathTodo)
	}

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "quiet",
				Usage:   "quiet the output",
				Value:   false,
				Aliases: []string{"q"},
			},
			&cli.StringFlag{
				Name:    "path",
				Usage:   "the path to your todo list",
				Value:   defaultPathTodo,
				Aliases: []string{"p"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:      "write",
				Aliases:   []string{"add", "touch"},
				Usage:     "write a new task to your todo list",
				ArgsUsage: "TASK",
				Action:    cmd.Write,
			},
			{
				Name:    "print",
				Aliases: []string{"list", "show", "ls"},
				Usage:   "read all the task in your todo list an put on the STDOUT",
				Action:  cmd.Print,
			},
			{
				Name:      "mark",
				Aliases:   []string{"do"},
				Usage:     "mark a task in your todo list as done",
				ArgsUsage: "TASK ID",
				Action:    cmd.Mark,
			},
			{
				Name:      "unmark",
				Aliases:   []string{"undo"},
				Usage:     "unmark an already done task in your todo list",
				ArgsUsage: "TASK ID",
				Action:    cmd.Unmark,
			},
			{
				Name:      "erase",
				Aliases:   []string{"delete", "rm", "purge", "remove"},
				Usage:     "delete a task in your todo list",
				ArgsUsage: "TASK ID",
				Action:    cmd.Erase,
			},
			{
				Name:    "brush",
				Aliases: []string{"clean", "clear"},
				Usage:   "remove forever all the tasks that you have done",
				Action:  cmd.Brush,
			},
		},
		Name:  "toddo",
		Usage: "simple CLI todo application",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
