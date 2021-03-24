package main

import (
	"fmt"

	"github.com/indeedhat/gli"
)

const (
	ErrOk = iota
	ErrPartial
	ErrBuggered
)

type PodCtl struct {
	Logs LogCommand `gli:"logs"`
}

// Run is the entry point called by the gli framework if the app is called without a sub command
// so far it does nothing
func (app *PodCtl) Run() int {
	fmt.Println("not implemented yet try 'podctl logs'")
	return ErrBuggered
}

func main() {
	app := gli.NewApplication(&PodCtl{}, "Podctl")
	app.Debug = true

	app.Run()
}
