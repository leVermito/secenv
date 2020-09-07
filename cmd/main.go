package main

import (
	"fmt"
	"github.com/Vermibus/secenv/internal/ciphers"
	"github.com/Vermibus/secenv/internal/environments"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	// encryptDecryptTest()
	console()
}

func encryptDecryptTest() {
	data := []byte("secret data")
	fmt.Printf("Secret: %s\n", data)

	key := []byte("abcdefghijklmnoprstwxyz123456789")
	nonce, cipherData := ciphers.EncryptAESGCM(key, data)

	fmt.Printf("len of cipherData: %d\n", len(cipherData))

	decryptedData := ciphers.DecryptAESGCM(nonce, key, cipherData)
	fmt.Printf("decryptedData: %s\n", decryptedData)
}

func console() {
	app := &cli.App{
		Name:     "secenv",
		HelpName: "secenv",
		Usage:    "Not quite insecure secret environments manager.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "env",
				Usage:    "Name of secret environment to manipulate.",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "path",
				Usage:       "Path to secenv directory.",
				DefaultText: "$HOME/.secenv",
				Required:    false,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "Create new secret environment.",
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
						Name:  "add, a",
						Usage: "Add new variable to secret environment.",
					},
					&cli.StringFlag{
						Name:  "edit, e",
						Usage: "Edit variable from secret environment.",
					},
					&cli.StringFlag{
						Name:  "remove, r",
						Usage: "Remove variable from secret environment.",
					},
				},
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:  "remove",
				Usage: "Remove secret environment.",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:  "inject",
				Usage: "Inject variables from secret environment to current session.",
				Action: func(c *cli.Context) error {
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
