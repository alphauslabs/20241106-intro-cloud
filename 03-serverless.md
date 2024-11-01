Deploy https://github.com/rastapasta/mapscii to Cloud Run.

``` sh
$ TBD
```

Deploy https://filebrowser.org/installation to Cloud Run.

``` sh
$ gcloud run deploy mapscii \
    --project=labs-169405 \
    --image=asia.gcr.io/labs-169405/mapscii:v4 \
    --region=asia-northeast1 \
    --allow-unauthenticated
```
