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
	Help bool        `gli:"^help,h" description:"Show this help document"`
	Logs LogCommand  `gli:"logs" description:"Follow logs from all pods matching the config"`
	List ListCommand `gli:"list" description:"List all pods matching the config"`
}

// Run is the entry point called by the gli framework if the app is called without a sub command
// so far it does nothing
func (app *PodCtl) Run() int {
	fmt.Println("not implemented yet try 'podctl logs'")
	return ErrBuggered
}

// NeedHelp defines if the command should end early and display the documentation
func (app *PodCtl) NeedHelp() bool {
	return app.Help
}

func main() {
	app := gli.NewApplication(&PodCtl{}, "Podctl")
	app.Debug = true

	app.Run()
}
