# CI building

### For the example we will use concourse CI

---

- Install concourse CI in minikube
```
k create ns ci
helm install --name my-release stable/concourse --namespace ci
```

- Get & expose the concourse URL
```
k -n ci expose pod/<my-release-web-*> -p 8080:8080
This should be: 127.0.0.1:8080
```

- run the following commands, to login, and push the pipeline then execute.
```
Login: fly --target golang login --concourse-url http://127.0.0.1:8080 -u test -p test

Push: fly -t ci set-pipeline -p golang -c pipeline.yaml

unpause pipeline: fly -t golang unpause-pipeline -p golang 
```
