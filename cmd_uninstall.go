package main

import (
    "fmt"

    "github.com/posener/complete/v2/install"
)

type UninstallCommand struct {
}

func (cmd *UninstallCommand) Run() int {
    if err := install.Uninstall("podctl"); err != nil {
        panic(err)
    }

    fmt.Println(`Autocomplete uninstalled!
You may need to source your shells rc file or restart your terminal for this to take effect`)

    return ErrOk
}
