#!/bin/bash

openapi-generator generate \
  -i ./openapi.yaml \
  -g go \
  -o client/ \
  --template-dir ./templates/go \
  --git-host github.com \
  --git-user-id mkcr-innovations \
  --git-repo-id xero-go-client/accounting/client \
  --additional-properties generateInterfaces=true,packageName=client \
  --global-property apis,models,supportingFiles=utils.go:configuration.go:client.go:README.md:go.mod:go.sum,apiTests=false
