package main

import (
	"github.com/chrisdoherty4/xavier/xavier/app/xavier"
	"github.com/chrisdoherty4/xavier/xavier/pkg/cmdutil"
)

func main() {
	cmdutil.Execute(xavier.Run)
}
