package main

import (
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
    Help      bool             `gli:"help,h" description:"Show this help message"`

	Apply     ApplyCommand     `gli:"apply" description:"Apply the changes to the pods kubernetes config"`
	Attach    AttachCommand    `gli:"attach" description:"Attach an interactive terminal to each matching pod"`
	Configure ConfigureCommand `gli:"configure" description:"Open up the kubernetes config files in your editor"`
	Exec      ExecCommand      `gli:"exec" description:"execute a comand on all pods that match"`
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
    return app.List.Run()
}

func (app *PodCtl) NeedHelp() bool {
    return app.Help
}

func main() {
    RegisterAutoComplete()

	app = gli.NewApplication(&PodCtl{}, `    ____            __     __  __  
   / __ \____  ____/ /____/ /_/ /  
  / /_/ / __ \/ __  / ___/ __/ /   
 / ____/ /_/ / /_/ / /__/ /_/ / 
/_/    \____/\__,_/\___/\__/_/  
    `)
	app.Debug = "" != os.Getenv("DEBUG")

	app.Run()
}
