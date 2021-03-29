package main

import (
	"path"
)

type ApplyCommand struct {
	conf *Config
}

// Run is the entry point to the apply command called by the gli framework
func (cmd *ApplyCommand) Run() int {
	var err error
	cmd.conf, err = loadConfig()

	if err != nil {
		panic(err)
	}

	applyCommand := append(
		CommandApply,
		path.Join(cmd.conf.Env.ConfigDir, cmd.conf.Pod.Name),
	)

	if _, err = kubectl(cmd.conf.Pod.Name, applyCommand); err != nil {
		panic(err)
	}

	return ErrOk
}
