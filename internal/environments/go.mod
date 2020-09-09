module github.com/Vermibus/secenv/internal/environments

replace github.com/Vermibus/secenv/internal/ciphers => ../ciphers

go 1.15

require (
	github.com/Vermibus/secenv/internal/ciphers v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
)
