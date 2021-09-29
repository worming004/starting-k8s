# az + kubectl 

this docker file build an image containing both azure cli, and kubectl cli. Both are used in this tutorial.

Build it with `docker build -t k8s-cli .`, then run it with `docker run -it k8s-cli`.

> WARNING: this you will run a container that will not be deleted after usage. It means that it would le re-use with the same Azure login. If you do not plan to re-use later this container, add `--rm` in the command.
