package main

import (
	"os"
	"os/exec"
	"path"
)

type ConfigureCommand struct {
	conf *Config
}

// Run is the entry point for the configure command called by the gli framework
func (cmd *ConfigureCommand) Run() int {
	var err error
	cmd.conf, err = loadConfig()

	if err != nil {
		panic(err)
	}

	editor := exec.Command(
		cmd.conf.Env.Editor,
		path.Join(cmd.conf.Env.ConfigDir, cmd.conf.Pod.Name),
	)

	editor.Stdin = os.Stdin
	editor.Stdout = os.Stdout
	editor.Stderr = os.Stderr

	err = editor.Run()

	if err != nil {
		panic(err)
	}

	return ErrOk
}
