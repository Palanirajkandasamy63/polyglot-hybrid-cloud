# DevOps Kubernetes Demo

This repository contains two microservices deployed to Kubernetes:
- **Python Flask App** - Simple web app with health checks
- **Go Web App** - Lightweight web service with detailed health endpoints

## Repository Structure
```
devops-k8s-demo/
├── python-app/          # Python Flask microservice
├── go-app/              # Go microservice
├── k8s/                 # Shared Kubernetes configs
└── scripts/             # Automation scripts
```

## Prerequisites

- Docker
- kind (Kubernetes in Docker)
- kubectl

## Quick Start
```bash
# 1. Create kind cluster
kind create cluster --name demo-cluster

# 2. Build both images
./scripts/build-all.sh

# 3. Deploy both apps
./scripts/deploy-all.sh

# 4. Access apps
kubectl port-forward service/python-flask-service 5000:80 &
kubectl port-forward service/go-app-service 8080:80 &

# Visit:
# http://localhost:5000 - Python app
# http://localhost:8080 - Go app
```

## Individual App Deployment

See individual READMEs:
- [Python App](./python-app/README.md)
- [Go App](./go-app/README.md)

## Cleanup
```bash
./scripts/cleanup.sh
```

## What You'll Learn

- Containerizing Python and Go applications
- Multi-stage Docker builds
- Kubernetes Deployments and Services
- Health checks (liveness and readiness probes)
- Resource management
- Scaling applications
- Rolling updates

## Author

Raj - DevOps Interview Preparation Project
