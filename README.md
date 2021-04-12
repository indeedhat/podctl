# podctl

## ISSUES
- [ ] the two teminal emulators i have tried (alacritty and xfce4-terminal) handle -c differently

## sub commands
- [x] logs `kubectl logs -f {{pod.name}}`
- [x] list `kubectl get pod | grep {{pod.name}}`
- [x] configure `$EDITOR {{env.config_dir}}`
- [x] apply `kubectl apply -f {{env.config_dir}}/{{pod.name}}`
- [x] restart  (i currently am not sure how im gonna do that)
- [x] exec `kubectl exec {{pod.id}} == {{arg[0]}}`
- [ ] attach `kubectl exec --tty --stdin {{pod.id}} -- {{pod.shell}}`
    - much of the work is done for this but i am having issues keeping terminals open with the attached shell
- [ ] info `kubectl describe {{pod.name}}`
- [ ] init (will prompt the user to enter details about the project | error if the config file exists)


## config file .podctl.toml
```toml
[pod]
name = "example-app"
# will defualt to default if not set
namespace = "some-namespace"
# she shell used by the container (defaults to sh)
shell = "bash"
# will default to kill $(top -n 1 | sed -n 5p | awk '{ print $1 }')
# to kill the main process, not graceful but seems to work
restart_cmd = ""

[env]
# will default to ~/.config/podctl if not set
config_dir = "/path/to/config/dir"
# will default to $EDITOR if not set
editor = "vim"
# will default to attempting to find a common terminal emulator on your system
terminal_emulator = "alacritty"


# logs is not yet implemented
[logs]
# defaults to "index" if not set
# can be one of the following:
# - index: will prefix each log line with the index of the pod as it appears in the kubectl get pod list
# - podId: will prefix each log with the full pod name/id as it appears int he kubectl get pod list
# - server-pod: will prefix each log with the server and pod suffix from kubectl get pod
# - server: will prefix each log with just the server suffix from kubectl get pod
# - pod: will prefix each log with just the pod suffix from kubectl get pod
# - none: log lines will not be prefixed
prefix = "index"
```
