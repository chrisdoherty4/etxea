package main

import (
	"github.com/chrisdoherty4/etxea/gela/pkg/app/gela"
	"github.com/chrisdoherty4/etxea/pkg/cmdutil"
)

func main() {
	cmdutil.Execute(gela.Run)
}
