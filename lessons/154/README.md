The purpose of this project is to simplify and automate the configuration of a Prometheus based monitoring stack for Kubernetes clusters.

kubectl apply --server-side


1. Create EKS cluster
terraform init
terraform apply
aws eks update-kubeconfig --name demo --region us-east-1

2. Install Prometheus Operator
kubectl apply --server-side -f prometheus-operator/crds
kubectl apply -f prometheus-operator/rbac/cluster-roles.yaml
kubectl apply -f prometheus-operator/namespace.yaml
kubectl apply -f prometheus-operator/deployment
kubectl get pods -n monitoring
kubectl logs -l app.kubernetes.io/name=prometheus-operator -n monitoring -f

2. Deploy Prometheus
kubectl apply -f prometheus
kubectl get pods -n monitoring
kubectl logs -l app.kubernetes.io/name=prometheus -n monitoring -f
kubectl get services -n monitoring
kubectl port-forward svc/prometheus-operated 9090 -n monitoring


3. Deploy Sample App

kubectl apply -f myapp/deploy
kubectl apply -f myapp/deploy/pod-monitor.yaml