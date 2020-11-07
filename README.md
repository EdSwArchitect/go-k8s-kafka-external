# Application in K8S Accessing Kafka, External

You'll have to set the IP address in kafka-reader-pod.yaml to the right IP address

## Order of K8S Command

1. minikube start --driver=[docker | virtualbox]
1. minikube kubectl -- apply -f kafka-external-svc.yaml
1. minikube kubectl -- apply -f kafka-reader-endpoint.yaml
1. minikube kubectl -- apply -f kafka-reader-pod.yaml
1. minikube kubectl -- apply -f kafka-reader-nodeport.yaml
1. minikube service list # Shows services and URL to access inside the pod for the metrics