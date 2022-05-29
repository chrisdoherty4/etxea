package main

import (
	"github.com/chrisdoherty4/xavier/internal/app/etxea"
	"github.com/chrisdoherty4/xavier/internal/pkg/cmdutil"
)

func main() {
	cmdutil.Execute(etxea.Run)
}
