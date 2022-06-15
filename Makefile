.PHONY: clean build build-all build-macos build-linux
.PHONY: generate lint fieldalign fieldalign-fix fmt
.PHONY: build-protos

#GO=$(shell which go)
GO=$(shell which gotip)
CMD_DIR=./cmd
BIN_DIR=./bin

MTADATA_CMD_DIR=$(CMD_DIR)/mtadata
MIGRATE_CMD_DIR=$(CMD_DIR)/migrate
PARSE_CMD_DIR=$(CMD_DIR)/parse
API_CMD_DIR=$(CMD_DIR)/api

MTADATA_BIN=mtadata
MIGRATE_BIN=migrate
PARSE_BIN=parse
API_BIN=api

GOOS_MACOS=darwin
GOOS_LINUX=linux

GOARCH_AMD64=amd64

MOD_NAME=github.com/jalavosus/mtadata

MODELS_PKG=$(MOD_NAME)/models

GENERATE_PKGS=$(MODELS_PKG)/boroughs $(MODELS_PKG)/divisions $(MODELS_PKG)/routes $(MODELS_PKG)/structures

PROTOS_DIR=./models/protos/v1/proto

define _build_os
@rm -rf $1
GOARCH=$(GOARCH_AMD64) GOOS=$3 $(GO) build -o $(BIN_DIR)/$3/$1 $2
endef

define _build
@rm -rf $(BIN_DIR)/$1
GOARCH=$(GOARCH_AMD64) GOOS=$3 $(GO) build -o $(BIN_DIR)/$1 $2
endef

define _fieldalign
fieldalignment $(strip $1) $2
endef


build :
	$(call _build,$(MTADATA_BIN),$(MTADATA_CMD_DIR),$(GOOS_MACOS))
	$(call _build,$(MIGRATE_BIN),$(MIGRATE_CMD_DIR),$(GOOS_MACOS))
	$(call _build,$(PARSE_BIN),$(PARSE_CMD_DIR),$(GOOS_MACOS))
	$(call _build,$(API_BIN),$(API_CMD_DIR),$(GOOS_MACOS))

build-all : clean build-macos build-linux

build-macos :
	$(call _build_os,$(MTADATA_BIN),$(MTADATA_CMD_DIR),$(GOOS_MACOS))
	$(call _build_os,$(MIGRATE_BIN),$(MIGRATE_CMD_DIR),$(GOOS_MACOS))
	$(call _build_os,$(PARSE_BIN),$(PARSE_CMD_DIR),$(GOOS_MACOS))
	$(call _build_os,$(API_BIN),$(API_CMD_DIR),$(GOOS_MACOS))

build-linux :
	$(call _build_os,$(MTADATA_BIN),$(MTADATA_CMD_DIR),$(GOOS_LINUX))
	$(call _build_os,$(MIGRATE_BIN),$(MIGRATE_CMD_DIR),$(GOOS_LINUX))
	$(call _build_os,$(PARSE_BIN),$(PARSE_CMD_DIR),$(GOOS_LINUX))
	$(call _build_os,$(API_BIN),$(API_CMD_DIR),$(GOOS_LINUX))

build-protos :
	buf build
	buf generate

clean :
	rm -rf $(BIN_DIR)
	mkdir -p $(BIN_DIR)/$(GOOS_MACOS) $(BIN_DIR)/$(GOOS_LINUX)

generate :
	go generate $(GENERATE_PKGS)

lint :
	@golangci-lint run --config=.golangci.yml

fieldalign :
	$(call _fieldalign,,$(FILES))

fieldalign-fix :
	$(call _fieldalign,-fix,$(FILES))

fmt :
	gofmt -s -w ./