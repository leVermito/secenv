package console

import (
	"github.com/Vermibus/secenv/internal/environments"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// Start : start command line parser
func Start() {
	app := &cli.App{
		Name:     "secenv",
		HelpName: "secenv",
		Usage:    "Not quite insecure secret environments manager.",
		Commands: []*cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create new secret environment.",
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
				Name:    "show",
				Aliases: []string{"s"},
				Usage:   "List variables in environment.",
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
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List environments.",
				Action: func(c *cli.Context) error {
					environments.ListEnvironments()
					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r"},
				Usage:   "Remove secret environment.",
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
				Name:    "inject",
				Aliases: []string{"i"},
				Usage:   "Inject variables from secret environment to current session.",
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
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add variable to secret environment.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "env",
						Usage:    "Name of secret environment to manipulate.",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					environments.AddVariableToEnvironment(c.String("env"))
					return nil
				},
			},
			{
				Name:    "edit",
				Aliases: []string{"e"},
				Usage:   "Edit variable from secret environment.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "env",
						Usage:    "Name of secret environment to manipulate.",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					environments.EditVariableFromEnvironment(c.String("env"))
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Delete variable from secret environment.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "env",
						Usage:    "Name of secret environment to manipulate.",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					environments.RemoveVariableFromEnvironment(c.String("env"))
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
