package cmdutil

import (
	"fmt"
	"os"
)

// Entrypoint is a command entrypoint.
type Entrypoint func(args []string) error

// Execute executes a command entrypoint.
func Execute(app Entrypoint) {
	if err := app(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
