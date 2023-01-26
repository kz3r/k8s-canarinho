# k8s-canarinho

**This is an intro Golang tryout to help me better understand the basics of the language.**

### What does it do

It's a simple app to check for any canaries running on my Kubernetes environment for longer that it should, and send a desktop notification if the conditions are met.

Limitations
- Kubernetes configurations are loaded based on the current active context set in kubeconfig
- Canary pod identification parameter is hardcoded

### Setup

The verification is based on namespaces. Those should listed under `namespaces` in the `conf.yaml` file, along with the maximum runtime allowed(in minutes) before the canaries are considered old and trigger the notification.

### Local build

<sub>requires golang 1.18</sub>

```bash
go build -o k8s-canarinho main.go
```