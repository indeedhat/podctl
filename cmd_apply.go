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
		fmt.Println("failed to load config")
		panic(err)
	}

	applyCommand := append(
		CommandApply,
		cmd.conf.Env.ConfigDir,
	)

	output, err := kubectl(cmd.conf.Pod.Namespace, applyCommand)
	fmt.Println(output)

	if err != nil {
		panic(err)
	}

	return ErrOk
}
