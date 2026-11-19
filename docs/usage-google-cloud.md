---
layout: page
title: Google Cloud Run
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Google Cloud Run

To get started in a hosted environment you can deploy this project to the Google Cloud Platform.

From your [Google Cloud dashboard](https://console.cloud.google.com/home/dashboard) create a new project and call it:
```
moov-icl-demo
```

Enable the [Artifact Registry](https://cloud.google.com/artifact-registry/docs/docker/quickstart) API for your project and associate a [billing account](https://cloud.google.com/billing/docs/how-to/manage-billing-account) if needed. Then, open the Cloud Shell terminal and run the following commands, substituting your unique project ID:

```
gcloud services enable artifactregistry.googleapis.com
gcloud artifacts repositories create imagecashletter \
  --repository-format=docker --location=us
gcloud auth configure-docker us-docker.pkg.dev

docker pull moov/imagecashletter
docker tag moov/imagecashletter us-docker.pkg.dev/<PROJECT-ID>/imagecashletter/imagecashletter
docker push us-docker.pkg.dev/<PROJECT-ID>/imagecashletter/imagecashletter
```

Deploy the container to Cloud Run:
```
gcloud run deploy --image us-docker.pkg.dev/<PROJECT-ID>/imagecashletter/imagecashletter --port 8083
```

Select your target platform to `1`, service name to `imagecashletter`, and region to the one closest to you (enable Google API service if a prompt appears). Upon a successful build you will be given a URL where the API has been deployed:

```
https://YOUR-ICL-APP-URL.a.run.app
```

Now you can list files stored in-memory:
```
curl https://YOUR-ICL-APP-URL.a.run.app/files
```
You should get this response:
```
null
```
