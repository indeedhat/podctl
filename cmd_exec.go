package main

import (
	"errors"
	"fmt"
	"os/exec"
	"sync"
)

type ExecCommand struct {
	Args []string

	conf *Config
	pods []*PodInfo

	wg     sync.WaitGroup
	errors map[string]error
}

func (cmd *ExecCommand) Run() int {
	var err error

	if len(cmd.Args) == 0 {
		panic(errors.New("No arguments given to exec"))
	}

	cmd.conf, err = loadConfig()

	if err != nil {
		panic(err)
	}

	cmd.pods, err = getPodIds(cmd.conf)

	if err != nil {
		panic(err)
	} else if len(cmd.pods) == 0 {
		panic("no pods to watch")
	}

	for i, info := range cmd.pods {
		cmd.exec(info, i)
	}

	cmd.wg.Wait()

	if len(cmd.errors) > 0 {
		if _, ok := cmd.errors["__skip"]; !ok {
			for index, err := range cmd.errors {
				fmt.Println(index, err)
			}
		}
		return ErrPartial

	}

	return ErrOk
}

func (cmd *ExecCommand) exec(info *PodInfo, i int) {
	cmd.wg.Add(1)

	go func() {
		terminal := exec.Command(
			cmd.conf.Env.TerminalEmulator,
			append(
				[]string{
					"-e",
					KubeCtl,
					"-n",
					cmd.conf.Pod.Namespace,
					"exec",
					info.PodID(),
					"--",
				},
				cmd.Args...,
			)...,
		)

		err := terminal.Run()

		if err != nil {
			index := info.Index(cmd.conf, i)

			if index == "" {
				fmt.Println(err)
				cmd.errors["__skip"] = err
			} else {
				cmd.errors[index] = err
			}
		}

		cmd.wg.Done()
	}()
}
