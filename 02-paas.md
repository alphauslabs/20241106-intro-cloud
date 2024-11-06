Use GKE for PaaS demo.

``` sh
# Create a zonal cluster:
$ gcloud container clusters create zz \
    --project=labs-169405 \
    --zone=asia-northeast1-b \
    --node-locations=asia-northeast1-b \
    --machine-type=e2-standard-2 \
    --scopes=default \
    --enable-autoscaling \
    --enable-vertical-pod-autoscaling \
    --min-nodes=1 \
    --max-nodes=2 \
    --maintenance-window=16:00 \
    --network=default \
    --subnetwork=default \
    --enable-ip-alias \
    --addons=HttpLoadBalancing \
    --release-channel=regular

# Point to cluster:
$ gcloud container clusters get-credentials zz --project labs-169405 --zone asia-northeast1-b

# Demo:
# 1) See logs using kubectl logs
# 2) Display number of nodes
# 3) Display number of pods
# 4) Initiate workload for scaling:
#    kubepfm --target simwork:8080:8080
#    curl -v -H "x-timeout-s: 30" localhost:8080/api
#    See if it scaled to multiple pods.
# 5) Revisit some time later and see scaling down.

# Delete cluster:
$ gcloud container clusters --project labs-169405 --zone asia-northeast1-b delete zz --quiet
```
