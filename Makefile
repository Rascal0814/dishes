GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
override MYSQL_DSN :="root:123456@tcp(127.0.0.1:3306)/order-dishes"
override MIGRATE_MYSQL_DSN := "mysql://"${MYSQL_DSN}
override MIGRATION_DIR :=./migrations

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

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

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