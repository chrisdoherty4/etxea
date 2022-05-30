package main

import (
	"github.com/chrisdoherty4/xavier/etxea/app/etxea"
	"github.com/chrisdoherty4/xavier/xavier/pkg/cmdutil"
)

func main() {
	cmdutil.Execute(etxea.Run)
}
