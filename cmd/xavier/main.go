package main

import (
	"github.com/chrisdoherty4/xavier/internal/app/xavier"
	"github.com/chrisdoherty4/xavier/internal/pkg/cmdutil"
)

func main() {
	cmdutil.Execute(xavier.Run)
}
