package main

import (
	"fmt"
	"strings"
	"sync"
)

type RestartCommand struct {
	conf *Config
	pods []*PodInfo

	wg sync.WaitGroup
}

func (cmd *RestartCommand) Run() int {
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

	for _, info := range cmd.pods {
		cmd.restart(info)
	}

	cmd.wg.Wait()

	return ErrOk
}

func (cmd *RestartCommand) restart(info *PodInfo) {
	cmd.wg.Add(1)

	if cmd.conf.Pod.RestartCmd != "" {
		cmd.customRestart(info)
		return
	}

	go func() {
		_, _ = kubectl(
			cmd.conf.Pod.Namespace,
			[]string{
				"exec",
				info.PodID(),
				"--",
				"kill",
				getMainProcessId(cmd.conf.Pod.Namespace, info.PodID()),
			},
		)

		cmd.wg.Done()
	}()
}

func (cmd *RestartCommand) customRestart(info *PodInfo) {
	go func() {
		_, _ = kubectl(
			cmd.conf.Pod.Namespace,
			[]string{
				"exec",
				info.PodID(),
				"--",
				cmd.conf.Pod.Shell,
				"-c",
				fmt.Sprintf("\"%s\"", cmd.conf.Pod.RestartCmd),
			},
		)
		cmd.wg.Done()
	}()
}

func getMainProcessId(namespace, podId string) string {
	out, err := kubectl(
		namespace,
		[]string{
			"exec",
			podId,
			"--",
			"top",
			"-n",
			"1",
		},
	)

	if err != nil || out == "" {
		return ""
	}

	lines := strings.Split(out, "\n")
	if len(lines) < 5 {
		return ""
	}

	return strings.Split(
		strings.Trim(lines[4], " "),
		" ",
	)[0]
}
