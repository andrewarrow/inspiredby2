gcloud auth login
gcloud config set project local-dev-353516
gcloud auth application-default login
brew install pulumi/tap/pulumi
pulumi new gcp-go
pulumi config set gcp:project local-dev-353516
pulumi up -s dev
pulumi destroy -s dev
gcloud compute instances list

