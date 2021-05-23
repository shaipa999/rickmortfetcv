# rickmortfetcv

 To Execute Script:
 
 
cd to fetcherscript folder and execute:
go run ./fetchrickandmort.go




 To Build docker:
 
 
cd to deploy folder and execute:
docker build -t rickmorthumali1:latest .



 To Start through kubernetes - reg yaml:
 
 
 cd to yamls folder and execute
 kubectl apply -f ./Deployment.yaml
 kubectl apply -f ./Service.yaml
 kubectl apply -f ./Ingress.yaml


 To Start through Helm:
 
 
 cd to helm folder and execute
 helm install rickmortprod rickmortyku



Fetch through:


 http://rickmorthumali1-ingress.com/fetch
 http://rickmorthumali1-ingress.com/healthcheck
