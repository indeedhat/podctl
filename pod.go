package main

import (
	"bufio"
	"strings"
)

const PodSuffixLen = 17

var CommandGetPod = []string{"get", "pod"}

// getPodIds will extract the podId's that match the pod name in config from the kubectl command
func getPodIds(conf *Config) (pods []string, err error) {
	stdout, err := kubectl(conf.Pod.Namespace, CommandGetPod)

	if err != nil {
		return
	}

	scanner := bufio.NewScanner(
		strings.NewReader(stdout),
	)

	// Skip the first line as that is headers
	scanner.Scan()

	for scanner.Scan() {
		if podId, ok := extractValidPodId(scanner.Text(), conf.Pod.Name); ok {
			pods = append(pods, podId)
		}
	}

	return
}

// extractValidPodId will compare the line against the podName given and only return a
// podId from the line if the check passes
//
// I am at this point not sure if the pod suffix is of fixed length but it will do for now
func extractValidPodId(line, podName string) (string, bool) {
	if !strings.HasPrefix(line, podName) {
		return "", false
	}

	podId := strings.Split(line, " ")[0]

	if len(podId) <= PodSuffixLen || podName != podId[:len(podId)-PodSuffixLen] {
		return "", false
	}

	return podId, true
}
