GCP VM:

``` sh
# Create:
$ gcloud compute instances create zz1 --machine-type e2-micro --zone asia-northeast1-a

# Delete:
$ gcloud compute instances delete zz1 --zone asia-northeast1-a --quiet
```
