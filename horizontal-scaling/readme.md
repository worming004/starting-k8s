# this folder demonstrate how horizontal pod scaling (HPA) is used by k8s

## deploy
build image as usually with `<your-docker-account>/horizontal-scaling` tag name
deploy service by replacing <your-docker-account> from deployment file.
deploy with `kubectl apply -f ./deployment.yaml -f service.yaml`

After few seconds, you will find the service IP by doing `kubectl get svc`. This is because we also declared a service resource, that request to Azure to bind an IP address

> NB: deployment describe that pod is allow to use 10% of a host cpu by precising `cpu: "100m"`. It's made in purpose to trigger easily an automated-scale based on CPU usage

## reach endpoints

`curl http://<ip-address>:8000/fib/30` to get fibonacci value for position 30

## apply automated scale

`kubectl autoscale deployment horizontal-scaling --cpu-percent=30 --min=1 --max=5`
this command autoscale the deployment *horizontal-scaling*. the min replicas number is 1, and max is 5. the Horizontal Pod Scaling (HPA) will try to balance number of replicates in order that cpu usage is near 30% by each replicas

## load test

run `go run ./loadtest/main.go -number 10 -address http://20.82.159.140:8000/fib/45` to run 10 parallel worker that will request the *address* in an infinite loop.

if you have not any go compiler installed, you can use docker with `docker run --rm -v ${PWD}/loadtest:/app golang go run /app/main.go -h`
