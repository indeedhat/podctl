package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

const (
	LineSeperator = "-----------------"
)

type ExecCommand struct {
	Args []string `gli:"!"`

	conf *Config
	pods []*PodInfo

	wg     sync.WaitGroup
	errors map[string]error
	stdout map[string]string
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

	cmd.errors = make(map[string]error)
	cmd.stdout = make(map[string]string)

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

	for index, err := range cmd.errors {
		fmt.Sprintf("%s\nError: %s\n%s\n", index, err, LineSeperator)
	}

	for index, out := range cmd.stdout {
		fmt.Sprintf("%s\n%s\n%s\n", index, out, LineSeperator)
	}

	if len(cmd.errors) > 0 {
		return ErrPartial

	}

	return ErrOk
}

func (cmd *ExecCommand) exec(info *PodInfo, i int) {
	cmd.wg.Add(1)

	go func() {
		log.Printf("pod %d", i)
		out, err := kubectl(
			cmd.conf.Pod.Namespace,
			append(
				[]string{
					info.PodID(),
					"--",
				},
				cmd.Args...,
			),
		)

		log.Println(out, err)
		index := info.Index(cmd.conf, i)
		if index == "" {
			index = fmt.Sprint(i)
		}

		cmd.errors[index] = err
		cmd.stdout[index] = out

		cmd.wg.Done()
	}()
}
