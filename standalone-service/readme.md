# deploy a standalone service

In this tutorial, we will learn how to deploy a service without any dependency. At the end, we will have an hello world service, accessible from internet.

## build the image

Build the image with `docker build -t <docker-account-name>/demo-standalone:latest` then push it with on your docker registry with `docker push <docker-account-name>/demo-standalone:latest`. By pushing it, your container is available on hub.docker.com, and accessible from your current k8s cluster for downlaod.

## adapt the deployment.yaml

The current deployment.yaml do not know what is you docker account name. Replace it in line 17.

## file explaination before deploying

The yaml file contains 2 main parts

The first one is a description of a Deployment resource. Deployment is responsible to manage deployment, and rollout.

The second part is a description of a Service. It's used to describe that pods from previous Deployment are exposed as a public service.

### vocabulary

Pod : unit of deployment. It can be a single container like this example, or multiple containers running together. (adding side-car for example)

## run the service

Apply both Deployment and Service with `kubectl apply -f /deployment.yaml`.

## test the service with curl

After few seconds, pods will be deployed, service will be configured and a public IP will be assigned.
Run `kubectl get service` until the standalone line show you an EXTERNAL-IP. Then simply do `curl <ip-address>:3000/pod`

## deeper understanding

Here is some commands and very brief informations of resources we play with.

### common commands

Here is a list of commands that will help you to understand better what is happening

* `kubectl get <resource-type>`
  * `kubectl get pod`
  * `kubectl get service`
  * `kubectl get deployment`
  * `kubectl get replicaset` # managed by Deployment resource in this example. Ignore it for now.

this command list you all resources of your type.

* `kubectl describe <resource-type> <resource-name>`
  * `kubectl describe

### pod

A pod is a unit of deployment. It's running 1 or more containers. In our example, each 3 replicas are running 1 container.

### replicaset

A replica set described the desired state of pod

### deployment

A deployment describe the desired state of replicaset.
As the control pipe look lîke this : Deployment -> ReplicaSet -> Pod, Deployment is able to describe desired state of pods.
Deployments usefull to controll rollout of services.

## delete a pod

```bash
kubectl get pod
kubectl delete pod
kubectl get pod
```

Even if we deleted a pod, this is one immediatly rebuild. As we have a deployment requesting 3 runnings pods, when one of them is killed, it's immediatly replaced

## get pod logs

get logs of a single pod with following command :

```bash
kubectl logs <pod-name>
```

Or you can use a selector filter with `-l`. As we described in *deployment.yaml* in lines 8-10, we applied a selector **app=standalone**. By doing this, you get logs of multiple pods.

```bash
kubectl logs -l <selector-filter>
kubectl logs -l app=standalone
```
