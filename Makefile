CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest
PROJECT_NAME=${PROJECT_NAME}


-include .env
POSTGRES_USER= postgres
POSTGRES_PASSWORD = postgres
POSTGRES_HOST =localhost
POSTGRES_PORT = 5432
POSTGRES_DATABASE = rabbitmq_task
POSTGRESQL_URL='postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable'

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

clear:
	rm -rf ${CURRENT_DIR}/bin/*

network:
	docker network create --driver=bridge ${NETWORK_NAME}

migrate-local-up:
	migrate -database ${POSTGRESQL_URL} -path storage/migrations up

migrate-local-down:
	migrate -database ${POSTGRESQL_URL} -path migrations down

migrate-local-down-last:
	migrate -database ${POSTGRESQL_URL} -path migrations down 1

create-new-migration:
	migrate create -ext sql -dir migrations -seq $(name)

mark-as-production-image:
	docker tag ${REGISTRY}/${APP}:${TAG} ${REGISTRY}/${APP}:production
	docker push ${REGISTRY}/${APP}:production

swag-init:
	swag init -g api/main.go -o api/docs

run-events-tests:
	cd ${CURRENT_DIR}/events/test && go test -v

run:
	go run cmd/main.go

.DEFAULT_GOAL:=run