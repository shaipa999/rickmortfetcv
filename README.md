# rickmortfetcv

 To Execute Script:
 
 ```
cd to fetcherscript folder and execute:
go run ./fetchrickandmort.go
```



 To Build docker:
 
 ```
cd to deploy folder and execute:
docker build -t rickmorthumali1:latest .
```


 To Start through kubernetes - reg yaml:
 
 ```
 cd to yamls folder and execute
 kubectl apply -f ./Deployment.yaml
 kubectl apply -f ./Service.yaml
 kubectl apply -f ./Ingress.yaml
```

 To Start through Helm:
 
 ```
 cd to helm folder and execute
 helm install rickmortprod rickmortyku
```


Fetch through:

```
* You will need to add the IP to the hosts file.

 http://rickmorthumali1-ingress.com/fetch
 http://rickmorthumali1-ingress.com/healthcheck
 ```
 
 
 WorkFlow:
 
```
The workflow:
1. Builds the app through docker.
2. Publishes it on Azure kubernetes http://20.81.102.53:8080
http://20.81.102.53:8080/fetch
http://20.81.102.53:8080/healthcheck

3. Performs tests.

```

 
