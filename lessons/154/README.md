Intro

- show prometheus labels and podmonitor/servicemonitor labels. How to match

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

4. Show metrics in Prometheus Explorer

tester_duration_seconds{quantile="0.99"}
tester_duration_seconds_count
rate(tester_duration_seconds_count[1m])

5. Remove pod monitor
kubectl delete -f myapp/deploy/pod-monitor.yaml
check that there now targets (config is empty)

6. Create servicemonitor
kubectl apply -f myapp/deploy/prom-service.yaml
kubectl get svc -n staging
kubectl get endpoints -n staging
kubectl describe endpoints myapp-prom -n staging

kubectl apply -f myapp/deploy/service-monitor.yaml

7. Deploy Grafana using Helm

helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm search repo grafana
helm search repo grafana
helm show values grafana/grafana --version 6.50.7
helm show values grafana/grafana --version 6.50.7 > grafana-values.yaml
create 11-helm-provider.tf
create 12-grafana-render.tf
terraform init
terraform apply
create 13-grafana.tf

helm list -n monitoring

kubectl get pods -n monitoring
kubectl get secrets -n monitoring
kubectl get secrets grafana -o yaml -n monitoring
echo "YWRtaW4=" | base64 -d
echo "RjlFYldsYlgzSkdhekYyd2dlMEdaYllldllMb2RIQW4wUmFhalp1Ug==" | base64 -d
F9EbWlbX3JGazF2wge0GZbYevYLodHAn0RaajZuR

kubectl get svc -n monitoring
kubectl port-forward svc/grafana 3000:80 -n monitoring

create data source
http://prometheus-operated.monitoring:9090

create dashboard
rate(tester_duration_seconds_count[1m])
{{path}}
requests per seconds
panel name: Traffic


## Additional scrape configs

create ec2 with node exporter
create 14-ec2.tf

create 15-prometheus-iam.tf

create additional-scrape-configs.yaml

update
  additionalScrapeConfigs:
    name: additional-scrape-configs
    key: prometheus-additional.yaml

update service account
  annotations:
    eks.amazonaws.com/role-arn: "arn:aws:iam::<acc-id>:role/prometheus-demo"


kubectl apply -f prometheus/service-account.yaml
kubectl apply -f prometheus/additional-scrape-configs.yaml
kubectl apply -f prometheus/prometheus.yaml
k delete pod prometheus-main-0 -n monitoring
kubectl port-forward svc/prometheus-operated 9090 -n monitoring

## Probe