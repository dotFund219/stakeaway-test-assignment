# Deploy the backend service on kubernetes

```
kubectl apply -f backend-config.yaml
kubectl apply -f backend-deployment.yaml
kubectl apply -f backend-service.yaml
kubectl apply -f prometheus-config.yaml
kubectl apply -f prometheus-deployment.yaml
kubectl apply -f prometheus-service.yaml
kubectl apply -f grafana-deployment.yaml
kubectl apply -f grafana-service.yaml
```

# Check Deployment Status

```
kubectl get pods
kubectl get services
```

# Access Monitoring Dashboards

- Prometheus: http://<node-ip>:9090
- Grafana: http://<node-ip>:3000
  - Default credentials: admin / admin
  - Add Prometheus as a data source.

# Create PromQL Queries

Use the following PromQL queries in Grafana:

```
sum by (endpoint) (rate(http_requests_total[5m]))
```

# Docker image link

dotfund/staking-backend-service:latest
