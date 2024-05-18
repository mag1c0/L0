include config.env
LOCAL_BIN:=$(CURDIR)/bin
export GOBIN ?= $(LOCAL_BIN)

install-deps:
	go install github.com/pressly/goose/v3/cmd/goose@latest

local-migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} down -v