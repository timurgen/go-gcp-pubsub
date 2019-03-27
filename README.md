# go-gcp-pubsub

Simple GCP pubsub sink for Sesam.io powered applications

env vars to configure service
* PORT - 8000 by default
* GOOGLE_APPLICATION_CREDENTIALS -path to credential files used by GCP lib, by default assuming file named credentials.json in work directory
* GOOGLE_APPLICATION_CREDENTIALS_CONTENT - GCP json authentication data as string
* GCP_PROJECT_ID - google cloud platform project id 

Service has one endpoint respoding to POST request with pub-sub topic name. POST body will be published to given pub sub topic.
