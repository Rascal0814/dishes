GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
override APP_NAME := dishes
override KUBE_CONTEXT := $(shell kubectl config current-context)
override KUBE_NAMESPACE := default
override MYSQL_DSN :="root:123456@tcp(43.143.80.123:3306)/order-dishes"
override MIGRATE_MYSQL_DSN := "mysql://"${MYSQL_DSN}
override MIGRATION_DIR :=./migrations
override GIT_REVISION := $(shell git rev-parse --short HEAD)
override BUILD_TIME := $(shell date '+%y%m%d')
override MIGRATION_OPTS := -database ${MIGRATE_MYSQL_DSN} -path $(MIGRATION_DIR)
ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

override DOCKER_REGISTRY := docker.io/0814l
override APP_IMAGE_NAME := $(DOCKER_REGISTRY)/$(APP_NAME)
override APP_IMAGE_TAG := $(BUILD_TIME)-$(GIT_REVISION)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install gorm.io/gen/tools/gentool@latest

# generate api proto
.PHONY: api
api:
	buf mod update && buf generate


.DEFAULT_GOAL := help

--check-feature:
FEATURE ?= $(shell bash -c 'read -p "Feature Name: " FEATURE; echo $$FEATURE')

gen-migration:--check-feature
	@migrate $(MIGRATION_OPTS) create -dir $(MIGRATION_DIR) -ext sql -seq $(FEATURE)

migrate-up:
	migrate $(MIGRATION_OPTS) -verbose up

migrate-reply:
	migrate $(MIGRATION_OPTS) -verbose down 1
	migrate $(MIGRATION_OPTS) -verbose up

gen-sql:
	gentool -dsn $(MYSQL_DSN) -db "mysql" -outPath "./internal/dao/query" -onlyModel

app:
	docker build -t $(APP_IMAGE_NAME):$(APP_IMAGE_TAG) -f Dockerfile .

push:
	docker push $(APP_IMAGE_NAME):$(APP_IMAGE_TAG)

deploy: app push
	@echo "Deploy by helm using context \033[0;31m$(KUBE_CONTEXT)\033[0m, waiting 5s to confirm"
	@sleep 5
	helm upgrade $(APP_NAME) devops/dishes \
		-n $(KUBE_NAMESPACE) \
		--set 'image.repository=$(APP_IMAGE_NAME)' \
		--set 'image.tag=$(APP_IMAGE_TAG)' \
		--install

.PHONY: mysql
mysql:
	@echo "Deploy by helm using context \033[0;31m$(KUBE_CONTEXT)\033[0m, waiting 5s to confirm"
	@sleep 5
	helm upgrade mysql devops/mysql \
		-n $(KUBE_NAMESPACE) \
		--install

d_mysql:
	docker run -itd -v /opt/mysql:/var/lib/mysql --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7

docker_build:
	ssh tx "docker rm -f dishes"
	ssh tx "docker images prune"
	ssh tx "docker run --restart=always -itd -p 8000:8000 \
			-v /root/app/order-dishes/config/local.yaml:/src/config/local.yaml \
 			--name dishes $(APP_IMAGE_NAME):$(APP_IMAGE_TAG)"

docker_deploy: app push docker_build
