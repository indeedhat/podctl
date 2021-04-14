package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var CommandGetPod = []string{"get", "pod"}

// PodInfo is the basic info about the pod... who knew
type PodInfo struct {
	Name         string
	ServerSuffix string
	PodSuffix    string
}

func (pod *PodInfo) PodID() string {
	return strings.Join([]string{
		pod.Name,
		pod.ServerSuffix,
		pod.PodSuffix,
	}, "-")
}

// index will build up the appropriate log prefix for displaying a log line
func (pod *PodInfo) Index(conf *Config, i int) string {
	var index string

	switch conf.Logs.Prefix {
	case LogIndexPodID:
		index = pod.PodID()

	case LogIndexServerPod:
		index = pod.ServerSuffix + "-" + pod.PodSuffix

	case LogIndexServer:
		index = pod.ServerSuffix

	case LogIndexPod:
		index = pod.PodSuffix

	default:
		index = strconv.Itoa(i)

	// yeah this is a little odd having default not be last but i wanted this to be a return
	case LogIndexNone:
		return ""
	}

	return fmt.Sprintf("[%s] - ", index)
}

// getPodIds will extract the podId's that match the pod name in config from the kubectl command
func getPodIds(conf *Config) (pods []*PodInfo, err error) {
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
		if info, ok := extractValidPodId(scanner.Text(), conf.Pod.Name); ok {
			pods = append(pods, info)
		}
	}

	return
}

// extractValidPodId will compare the line against the podName given and only return a
// podId from the line if the check passes
//
// I am at this point not sure if the pod suffix is of fixed length but it will do for now
func extractValidPodId(line, podName string) (*PodInfo, bool) {
	if !strings.HasPrefix(line, podName) {
		return nil, false
	}

	podId := strings.Split(line, " ")[0]
	parts := strings.Split(podId, "-")
	foundName := strings.Join(parts[:len(parts)-2], "-")

	if len(parts) < 3 || foundName != podName {
		return nil, false
	}

	info := &PodInfo{
		Name:         podName,
		ServerSuffix: parts[len(parts)-2],
		PodSuffix:    parts[len(parts)-1],
	}

	return info, true
}
