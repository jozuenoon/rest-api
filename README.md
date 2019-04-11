# FinTech API

Implements payments API.

Server works on port `8080`

# TODO

* add middleware for authorization, authentication,
* add input validation eg. with https://github.com/mwitkow/go-proto-validators,
* add proper authentication and authorization,
* make swagger documentation better, this requires writing annotations in proto file unfortunately, since only `http` annotation is supported with yaml files,
* add Dockerfile,
* write helm chart for K8S deployment,
* configure TLS if it's not terminated on LB or Ingress,