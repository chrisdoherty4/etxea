package main

import (
	"github.com/chrisdoherty4/etxea/pkg/cmdutil"
	"github.com/chrisdoherty4/etxea/pkg/etxea"
)

func main() {
	cmdutil.Execute(etxea.Run)
}
