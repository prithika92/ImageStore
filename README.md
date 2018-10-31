# ImageStore
#REST Service created using GoLang
StoreImage.go

#Docker Image created
Dockerfile

#To start Minikube
minikube start
#To open Docker Terminal
minikube ssh
#To copy the REST Service and Docker Image to the current directory in Docker from the Windows Machine
scp //c/Users/docker/* .

#Docker Build
docker build -t app:v1 .

#Deploy app
kubectl create -f deployment.yaml  

# Deploy service
kubectl create -f service.yaml 
