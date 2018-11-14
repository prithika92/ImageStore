# ImageStore to upload, list and delete images
#REST Service created using GoLang
StoreImage.go

#Docker file
Dockerfile

#Start the minikube and ssh to the minikube terminal
minikube start
minikube ssh
#To copy the REST Service and Docker file to the current directory in Docker from the Windows Machine
scp //c/Users/docker/* .

#Docker Build
docker build -t app:v1 .

#Create a persistant Volume and claim 
kubectl create -f volume.yaml  
kubectl create -f claim.yaml  


#Deploy app
kubectl create -f deployment.yaml  

# Deploy service
kubectl create -f service.yaml 

# To check if the created pod is up and running
kubectl get pod

# Get the minikube service url for the created service
minikube service image-store –-url
