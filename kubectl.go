package main

import (
	"bufio"
	"bytes"
	"errors"
	"log"
	"os/exec"
	"strings"
)

const (
	KubeCtl = "kubectl"
)

var (
	CommandFollowLog = []string{"logs", "-f", "-n"}
	CommandApply     = []string{"apply", "-f"}
)

// kubectl will run a kubectl command and return the error/output produced
func kubectl(namespace string, args []string) (string, error) {
	var stdout, stderr bytes.Buffer
	finalArgs := append(
		[]string{"-n", namespace},
		args...,
	)

	cmd := exec.Command(KubeCtl, finalArgs...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	log.Println(strings.Join(finalArgs, " "))
	if err := cmd.Run(); err != nil {
		log.Println("error returned from cmd")
		return stdout.String(), err
	}

	if stderr.Len() != 0 {
		log.Println("stderr is not empty")
		return stdout.String(), errors.New(stderr.String())
	}

	if stdout.Len() == 0 {
		log.Println("no stdout data")
		return stdout.String(), errors.New("command did not return any output")
	}

	return stdout.String(), nil
}

// kubectlFollow will setup a command to follow the logs from kubectl logs -f
func kubectlFollowLog(namespace, podId string) (*exec.Cmd, *bufio.Reader, error) {
	finalArgs := append(
		CommandFollowLog,
		namespace,
		podId,
	)

	log.Println(finalArgs)

	cmd := exec.Command(KubeCtl, finalArgs...)

	// soooo... yeah this is not stdout but stderr, for some reason kubectl passes logs to stderr when following
	// it actually seems that quite a few commands to this, maybe there is a reason... perhaps one day ill look into it
	// but for now I'm using it as stdout
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, nil, err
	}

	return cmd, bufio.NewReader(stdout), nil
}
