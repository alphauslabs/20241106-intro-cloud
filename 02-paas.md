Use GKE for PaaS demo.

``` sh
# Create cluster:
gcloud container clusters create zz \
       --project=mobingi-main \
       --zone=asia-northeast1-b \
       --machine-type=e2-standard-2 \
       --scopes=default \
       --enable-autoscaling \
       --enable-vertical-pod-autoscaling \
       --min-nodes=1 \
       --max-nodes=4 \
       --maintenance-window=16:00 \
       --network=default \
       --subnetwork=default \
       --enable-ip-alias \
       --addons=HttpLoadBalancing \
       --release-channel=regular

# Delete cluster:
$ gcloud container clusters --zone asia-northeast1-b delete zz --quiet
```
