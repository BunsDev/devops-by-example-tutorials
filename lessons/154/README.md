The purpose of this project is to simplify and automate the configuration of a Prometheus based monitoring stack for Kubernetes clusters.

kubectl apply --server-side


1. Create EKS cluster
terraform init
terraform apply
aws eks update-kubeconfig --name demo --region us-east-1

2. Install Prometheus Operator
kubectl apply --server-side -f prometheus-operator/crds
kubectl apply -f prometheus-operator/rbac/cluster-roles.yaml