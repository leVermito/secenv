package main

import (
	// "fmt"
	"github.com/Vermibus/secenv/internal/environments"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	console()
}

func console() {
	app := &cli.App{
		Name:     "secenv",
		HelpName: "secenv",
		Usage:    "Not quite insecure secret environments manager.",
		Flags:    []cli.Flag{
			// &cli.StringFlag{
			// 	Name:        "path",
			// 	Usage:       "Path to secenv directory.",
			// 	DefaultText: "$HOME/.secenv",
			// 	Required:    false,
			// },
		},
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "Create new secret environment.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "env",
						Usage:    "Name of secret environment to manipulate.",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					environments.CreateEnvironment(c.String("env"))
					return nil
				},
			},
			{
				Name:  "edit",
				Usage: "Edit secret environment",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "env",
						Usage:    "Name of secret environment to manipulate.",
						Required: true,
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:    "add",
						Aliases: []string{"a"},
						Usage:   "Add new variable to secret environment.",
						Action: func(c *cli.Context) error {
							environments.AddVariableToEnvironment(c.String("env"))
							return nil
						},
					},
					{
						Name:    "edit",
						Aliases: []string{"e"},
						Usage:   "Edit variable from secret environment.",
						Action: func(c *cli.Context) error {
							environments.EditVariableFromEnvironment(c.String("env"))
							return nil
						},
					},
					{
						Name:    "remove",
						Aliases: []string{"r"},
						Usage:   "Remove variable from secret environment.",
						Action: func(c *cli.Context) error {
							environments.RemoveVariableFromEnvironment(c.String("env"))
							return nil
						},
					},
				},
			},
			{
				Name:  "show",
				Usage: "List variables in environment.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "env",
						Usage:    "Name of secret environment to manipulate.",
						Required: true,
					},
					&cli.BoolFlag{
						Name:  "values",
						Usage: "Print values.",
					},
				},
				Action: func(c *cli.Context) error {
					environments.ShowEnvironment(c.String("env"), c.Bool("values"))
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "List environments.",
				Action: func(c *cli.Context) error {
					environments.ListEnvironments()
					return nil
				},
			},
			{
				Name:  "remove",
				Usage: "Remove secret environment.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "env",
						Usage:    "Name of secret environment to manipulate.",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					environments.RemoveEnvironment(c.String("env"))
					return nil
				},
			},
			{
				Name:  "inject",
				Usage: "Inject variables from secret environment to current session.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "env",
						Usage:    "Name of secret environment to manipulate.",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					environments.InjectVariablesFromEnvironment(c.String("env"))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
