# mario-friends

This is a test repository for a little bit of everything. It has configurations
for building the app mario-friends, and tasks for deploying it to a KIND
cluster

## Stuff used in this repository

- [Docker](https://www.docker.com) (or [podman](https://podman.io), support is
  experimental and is still a bit unstable with KIND)
- [KIND](https://kind.sigs.k8s.io/)
- [timoni](https://timoni.sh)
- [CUE](https://cuelang.org)
- [taskfile](https://taskfile.dev)

## Tasks

The following tasks have been defined in taskfile:

```sh
❯ task
Welcome to mario-friends
task: Available tasks for this project:
* create-cluster:             Create a new cluster
* deploy:                     Deploy the application to the cluster
* destroy-cluster:            Destroy cluster
* install-cert-manager:       Install cert-manager
* install-ingress:            Install nginx ingress
* kind-image:                 Build and upload a new image to the cluster
```

## Deploy

To get started just run `task deploy`. This will create a new KIND cluster,
install an nginx ingress operator, build a docker image with the mario-friends
application, upload the new image to the internal KIND cluster registry, and
deploy the current configuration.

The current configuration starts 3 pods: troopa, bowser and princess. When
you do a GET request to troopa, the troopa pod does a call to the bowser pod,
which then does a call to the princess pod. When everything is setup, the chain
produce this output:

```sh
curl localhost/mario-friends/troopa
You've passed troopa! 🎉
🎮 Let the games begin 🎮

Entering level bowser...


You've defeated bowser!
You are doing great!😎

Entering level princess...


You've saved the princess!
🤩 🤩 🤩 🤩 🤩 🤩 🤩 🤩 🤩
```
