package main

import (
	"fmt"
	"sync"
)

// LogCommand is the gli struct for the podctl logs command
// it is in charge of watching the logs for each pod that matches the given name
// and concatinating them into a single stream
type LogCommand struct {
	conf *Config
	pods []string

	wg     sync.WaitGroup
	errors []string
}

// Run is the entry point to the command called by the gli framework
func (cmd *LogCommand) Run() int {
	var err error
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

	for i, podId := range cmd.pods {
		cmd.follow(podId, i)
	}

	fmt.Println("before")
	cmd.wg.Wait()
	fmt.Println("after")

	if len(cmd.errors) > 0 {
		for _, err := range cmd.errors {
			fmt.Println("ERROR", err)
		}

		return ErrPartial
	}

	return ErrOk
}

// follow is designed to be ran as a goroutine for each pod that is being watched
func (cmd *LogCommand) follow(podId string, i int) {
	fmt.Println("adding", i)
	cmd.wg.Add(1)

	go func() {
		_, reader, err := kubectlFollowLog(cmd.conf.Pod.Namespace, podId)

		if err != nil {
			cmd.errors = append(cmd.errors, fmt.Sprintf("[%d] - %s", i, err.Error()))
			fmt.Println("error", i)
			cmd.wg.Done()
			return
		}

		for {
			line, _, err := reader.ReadLine()

			if err != nil {
				cmd.errors = append(cmd.errors, fmt.Sprintf("[%d] - %s", i, err.Error()))
				break
			}

			fmt.Printf("[%d] - %s\n", i, line)
		}

		fmt.Println("done", i)
		cmd.wg.Done()
	}()
}
