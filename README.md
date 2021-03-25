# podctl

## sub commands
- [.]list `kubectl get pod | grep {{pod.name}}`
    - functional but needs to have the prefix settings implemented
- [ ] init (will prompt the user to enter details about the project | error if the config file exists)
- [ ] logs `kubectl logs -f {{pod.name}}`
- [ ] restart  (i currently am not sure how im gonna do that)
- [ ] configure `$EDITOR {{env.config_dir}}`
- [ ] apply `kubectl apply -f {{env.config_dir}}/{{pod.name}}`
- [ ] info `kubectl describe {{pod.name}}`

## config file .podctl.toml
```toml
[pod]
name = "example-app"
# will defualt to default if not set
namespace = "some-namespace"

[env]
# will default to ~/.config/podctl if not set
config_dir = "/path/to/config/dir"
# will default to $EDITOR if not set
editor = "vim"

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
