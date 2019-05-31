Set the following environment variables:

```bash
export PROJECT_ID="your-project-id"
export KO_DOCKER_REPO="gcr.io/${PROJECT_ID}"
```

Then run:

```bash
ko apply -f k8s/app.yaml
```
