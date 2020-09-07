module github.com/Vermibus/secenv/cmd

go 1.15

replace github.com/Vermibus/secenv/internal/ciphers => ../internal/ciphers

replace github.com/Vermibus/secenv/internal/environments => ../internal/environments

require (
	github.com/Vermibus/secenv/internal/ciphers v0.0.0-00010101000000-000000000000
	github.com/Vermibus/secenv/internal/environments v0.0.0-00010101000000-000000000000
	github.com/urfave/cli/v2 v2.2.0
)
