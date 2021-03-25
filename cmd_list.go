package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ListCommand struct {
	conf *Config
}

// Run is the entry point to the list command called by the gli framework
func (cmd *ListCommand) Run() int {
	var err error
	cmd.conf, err = loadConfig()

	if err != nil {
		panic(err)
	}

	getPods, err := kubectl(cmd.conf.Pod.Namespace, CommandGetPod)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(
		strings.NewReader(getPods),
	)

	scanner.Scan()
	fmt.Println(scanner.Text())

	for scanner.Scan() {
		if _, ok := extractValidPodId(scanner.Text(), cmd.conf.Pod.Name); ok {
			fmt.Println(scanner.Text())
		}
	}

	return ErrOk
}
