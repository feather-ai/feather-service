#!/bin/bash
# Code generation

SHELL=/bin/bash

swagger:
	@swagger validate api/swagger.yml
	@mkdir -p api/generated
	@swagger generate server -A core -t api/generated -f api/swagger.yml --exclude-main

mocks:
	@mockgen -source=./lib/aisystems/aisystems.go -destination=./lib/aisystems/aisystemsmock/aisystems_mock.go -package aisystemsmock
	@mockgen -source=./lib/gatekeeper/gatekeeper.go -destination=./lib/gatekeeper/gatekeepermock/gatekeeper_mock.go -package gatekeepermock
	@mockgen -source=./lib/storage/storage.go -destination=./lib/storage/storagemock/storage_mock.go -package storagemock
	@mockgen -source=./lib/upload/upload.go -destination=./lib/upload/uploadmock/upload_mock.go -package uploadmock
	@mockgen -source=./lib/users/users.go -destination=./lib/users/usersmock/users_mock.go -package usersmock

# Build stages

build: 
	@CGO_ENABLED=0 go build cmd/service-core.go

run:
	@PORT=8080 AWS_REGION=us-east-2 AWS_PROFILE=featherai DEBUG_USER=true go run cmd/service-core.go

tests: 
	@go test ./lib/...

integration_tests:
	@go test ./lib/... -tags=integration
	
# DB management

migrate:
	@tern migrate --config ./db/tern.local.conf --migrations ./db/migrations/

migrate_down:
	@tern migrate --destination -1 --config ./db/tern.local.conf --migrations ./db/migrations/

docker:
	@docker build . -t ftr-service-core

# Cloud Environment

deploy: docker
	aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 483384053975.dkr.ecr.us-east-2.amazonaws.com
	docker tag ftr-service-core:latest 483384053975.dkr.ecr.us-east-2.amazonaws.com/ftr-service-core:latest
	docker push 483384053975.dkr.ecr.us-east-2.amazonaws.com/ftr-service-core:latest

# Run the service locally, but connect to the cloud DB
run_cloud:
	. .local/cloud.env && PORT=8080 AWS_REGION=us-east-2 DEBUG_USER=true go run cmd/service-core.go

migrate_cloud:
	. .local/cloud.env && tern migrate --config ./db/tern.cloud.conf --migrations ./db/migrations/

migrate_down_cloud:
	. .local/cloud.env && tern migrate --destination -1 --config ./db/tern.cloud.conf --migrations ./db/migrations/
