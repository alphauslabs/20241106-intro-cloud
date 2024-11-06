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

# Delete cluster:
$ gcloud container clusters --project labs-169405 --zone asia-northeast1-b delete zz --quiet
```
