package main

import "fmt"

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
		cmd.conf.Env.ConfigDir,
	)

	if output, err := kubectl(cmd.conf.Pod.Name, applyCommand); err != nil {
		fmt.Println(output)
		panic(err)
	}

	return ErrOk
}
