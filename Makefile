.PHONY: build-all lint fieldalign fieldalign-fix fmt

GO=$(shell which go)
CMD_DIR=./cmd
BIN_DIR=./bin

MTADATA_CMD_DIR=$(CMD_DIR)/mtadata
MIGRATE_CMD_DIR=$(CMD_DIR)/migrate
PARSE_CMD_DIR=$(CMD_DIR)/parse

MTADATA_BIN=$(BIN_DIR)/mtadata
MIGRATE_BIN=$(BIN_DIR)/migrate
PARSE_BIN=$(BIN_DIR)/parse

define _build
@rm -rf $1
$(GO) build -o $1 $2
endef

define _fieldalign
fieldalignment $(strip $1) $2
endef


build-all :
	$(call _build,$(MTADATA_BIN),$(MTADATA_CMD_DIR))
	$(call _build,$(MIGRATE_BIN),$(MIGRATE_CMD_DIR))
	$(call _build,$(PARSE_BIN),$(PARSE_CMD_DIR))

lint :
	@golangci-lint run --config=.golangci.yml

fieldalign :
	$(call _fieldalign,,$(FILES))

fieldalign-fix :
	$(call _fieldalign,-fix,$(FILES))

fmt :
	gofmt -s -w ./