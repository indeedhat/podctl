package main

import (
	"fmt"
	"os"
	"os/exec"
)

type ConfigureCommand struct {
	conf *Config

    Help  bool `gli:"help,h" description:"Show this message"`
    Print bool `gli:"print,p" description:"Print the config path"`
    Apply bool `gli:"apply,a" description:"Auto apply the configuration on editor colse\n    This will blindly apply as soon as the editor process finishes"`
}


// Run is the entry point for the configure command called by the gli framework
func (cmd *ConfigureCommand) Run() int {
	var err error
	cmd.conf, err = loadConfig()

	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(cmd.conf.Env.ConfigDir); os.IsNotExist(err) {
		if err := os.MkdirAll(cmd.conf.Env.ConfigDir, 0755); err != nil {
			panic(err)
		}
	}

    if cmd.Print {
        fmt.Print(cmd.conf.Env.ConfigDir)
        return ErrOk
    }

	editor := exec.Command(cmd.conf.Env.Editor)
	editor.Dir = cmd.conf.Env.ConfigDir

	editor.Stdin = os.Stdin
	editor.Stdout = os.Stdout
	editor.Stderr = os.Stderr

	err = editor.Run()

	if err != nil {
		panic(err)
	}

    if cmd.Apply {
        if err := applyConfig(cmd.conf); err != nil {
            panic(err)
        }
    }

	return ErrOk
}

func (cmd *ConfigureCommand) NeedHelp() bool {
    return cmd.Help
}
