```sh
VERSION=v1 skaffold build --quiet > build.out ; skaffold deploy --build-artifacts build.out
VERSION=v2 skaffold build --quiet > build.out ; skaffold deploy --build-artifacts build.out
```

```sh
export GATEWAY_IP=$(kubectl -n istio-system get svc/istio-ingressgateway -ojson \
    | jq -r .status.loadBalancer.ingress[0].ip)
echo $GATEWAY_IP
```

```sh
watch "kubectl describe canaries.flagger.app simple-web-app | tail -n 30"
```
