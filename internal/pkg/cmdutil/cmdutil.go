package cmdutil

import (
	"fmt"
	"os"
)

// Entrypoint is a command entrypoint.
type Entrypoint func(args []string) error

// Execute  a command entrypoint. If app returns an error, Execute prints the error and exits the
// program with a non-zero code.
func Execute(app Entrypoint) {
	if err := app(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err.Error())
		os.Exit(1)
	}
}
