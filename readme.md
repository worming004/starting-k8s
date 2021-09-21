# Introduction to kubernetes

This project describe how to create a small k8s cluster on azure, and how to deploy a service to internet.
In this tutorial, we will use as much as possible docker to not have to install anything else.

Also, this tutorial is planned to have something running in a demo presentation of few hours max. It will not explain every components in details. If you are interested to have a deep understanding, I highly recommend you [Kubernetes in action](https://www.manning.com/books/kubernetes-in-action)

Finally, this tutorial is not about Docker. Having a base understanding of what are images, containers, and be able to run an image locally is required.

## Prerequisite

* docker installed locally, using linux container. A base knowledge is required.
* having a azure account with a valid subscription. Take note of the id
* a free account in hub.docker.com
* optionnal : vscode with https://marketplace.visualstudio.com/items?itemName=ms-kubernetes-tools.vscode-kubernetes-tools extension

## Start a new cluster

### Cli

#### Build cli container

If you do not wants to install az cli & kubectl cli, you can use this cli to run everything you will need.

From build **buildcli** folder, build a cli container with `docker build -t demo-cli .`. From your current terminal, come back to the root *starting-l8s* directory, then access it with `docker run -it -v $(PWD):/workdir --rm demo-cli`. (from bash, replace $(PWD) by ${PWD}).

From container, run `source /usr/share/bash-completion/bash_completion && source <(kubectl completion bash)` to add autocompletion for kubectl command

#### Login

Run an azure cli from docker with `docker run -it --rm mcr.microsoft.com/azure-cli`. Now your current terminal is using a bash with azure cli installed. All commands are invoked with `az`.

Before doing more actions, you need to login to your account, with `az login`. Follow the instruction by opening url with your browser, and login to your account.

You should be ready to run any azure commands now.

### Create a resource group and the cluster

First, you need to create a resource group that will handle your cluster. Simply run `az group create --location <location> --name training`.

The location list can be found by double tapping *tab* when you already entered `az group create --location ` with white space. If you do not know what to put, you can select *westeurope*

If you already have a resource group named *training*, feel free to replace it with another name. We will use *training* name in all next steps of this tutorial.

Then, you are ready to create your cluster with `az aks create -g training -n default-cluster -s Standard_A2_v2 -c 2 --generate-ssh-keys`

From now, after a few times, the kubernetes cluster is ready to use !

> From now, if you want to **stop** the tutorial to not pay anymore for a small cluster, you can run `az aks stop -g trining -n default-cluster`. You will be able to restart later with `az aks start -g training -n default-cluster`
> If you are finish with this tutorial and do not plan to come back, you can **delete** it with `az aks delete -g training -n default-cluster`

## next tutorial

I am expecting you to follow the tutorial in the following order.

* ./standalone-service
* ./next if people enjoyed it ..
