.PHONY: vuln-scan env-vault

export VAULT_HOST ?= http://127.0.0.1:8200
export VAULT_TOKEN ?= token
export ENVCONSUL_CONFIG ?= ./config.hcl
export ENVCONSUL_SECRET_PATH ?= kv/secret
export CI_JOB_TOKEN ?= glpat-xxxxxx

vuln-scan:
	trivy fs --scanners vuln .

env-vault:
	echo "./envconsul -config="$(ENVCONSUL_CONFIG)" -vault-addr="$(VAULT_HOST)" -once ./executable" > cmd.sh

build-api:
	docker build -f deployments/api/Dockerfile --build-arg="CI_JOB_TOKEN=$(CI_JOB_TOKEN)" .

run-api:
	gofmt -w . && go run ./cmd/api/main.go

run-api-graphql:
	gofmt -w . && go run ./cmd/api-graphql/main.go

run-api-grpc:
	gofmt -w . && go run ./cmd/api-grpc/main.go

run-migration:
	gofmt -w . && go run ./cmd/migration/main.go

swagger-api:
	cd ./cmd/api && swag init --dir ./,../../internal/api-admin --output ../../internal/api-admin/docs --outputTypes go --parseDependency
