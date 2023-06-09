# Signals

Kubernetes controllers or deployments, or basically anything you run in Kubernetes, runs on the Pod, hence to understand how to do graceful termination- one needs to look at how pod termination is handled. Have a look at the Kuberntes documentation on termination of pods [here](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination).

For a rough summary, the container runtime sends a TERM signal to the main process in each container, and gives it a grace period (by default: 30 sec) to terminate failing which it sends the KILL signal. Note, however, that during development- you are more likely to press `ctrl + c` which sends a INT signal.

All of this is relatively easy to handle in Go. Have a look here for [ref](https://gobyexample.com/signals). In addition, we pass this down into the program via a context to cancel operations deeper down the stack. The Kubernetes libraries also accept context for this very purpose.