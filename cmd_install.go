package main

import (
	"fmt"

	"github.com/posener/complete/v2/install"
)

type InstallCommand struct {
}

func (cmd *InstallCommand) Run() int {
    if err := install.Install("podctl"); err != nil {
        panic(err)
    }

    fmt.Println(`Autocomplete installed!
You may need to source your shells rc file or restart your terminal for this to take effect`)

    return ErrOk
}
