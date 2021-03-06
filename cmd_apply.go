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

    if err  = applyConfig(cmd.conf); err != nil {
		panic(err)
	}

	return ErrOk
}

// applyConfig will run kubectl apply -f {config dir}
func applyConfig(conf *Config) error {
	applyCommand := append(
		CommandApply,
		conf.Env.ConfigDir,
	)

	output, err := kubectl(conf.Pod.Namespace, applyCommand)
	fmt.Println(output)

    return err
}
