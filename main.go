package main

import (
	"fmt"
	"os"

	"github.com/indeedhat/gli"
)

const (
	ErrOk = iota
	ErrPartial
	ErrBuggered
)

var (
    app *gli.App
)

type PodCtl struct {
	Apply     ApplyCommand     `gli:"apply" description:"Apply the changes to the pods kubernetes config"`
	Attach    AttachCommand    `gli:"attach" description:"Attach an interactive terminal to each matching pod"`
	Configure ConfigureCommand `gli:"configure" description:"Open up the kubernetes config files in your editor"`
	Exec      ExecCommand      `gli:"exec" description:"execute a comand on all pods that match"`
    Help      HelpCommand      `gli:"help" description:"Show this help message"`
    Init      InitCommand      `gli:"init" description:"Initialise the podctl environment for this directory"`
    Install   InstallCommand   `gli:"install" description:"install auto complete rules for podctl in your shell"`
	List      ListCommand      `gli:"list" description:"List all pods matching the config"`
	Logs      LogCommand       `gli:"logs" description:"Follow logs from all pods matching the config"`
	Restart   RestartCommand   `gli:"restart" description:"kill/restart all pods that match\n    In its default state it will only crash the pod and let k8s pick it back up again"`
    Uninstall UninstallCommand `gli:"uninstall" description:"uninstall auto complete rules for podctl in your shell"`
}

// Run is the entry point called by the gli framework if the app is called without a sub command
// so far it does nothing
func (app *PodCtl) Run() int {
	fmt.Println("not implemented yet try 'podctl logs'")
	return ErrBuggered
}

func main() {
    RegisterAutoComplete()

	app = gli.NewApplication(&PodCtl{}, "Podctl")
	app.Debug = "" != os.Getenv("DEBUG")

	app.Run()
}
