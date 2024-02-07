.PHONY: all test

all: build

build: protogen
	go build ./...

# {{{ Protobuf

# Protobuf definitions
PROTO_FILES := $(shell find . \( -path "./languages" -o -path "./specification" \) -prune -o -type f -name '*.proto' -print)
# Protobuf Go files
PROTO_GEN_FILES = $(patsubst %.proto, %.pb.go, $(PROTO_FILES))

# Protobuf generator
PROTO_GO_MAKER := protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative


protogen: $(PROTO_GEN_FILES)

%.pb.go: %.proto
	cd $(dir $<); $(PROTO_GO_MAKER) --proto_path=. --proto_path=$(GOPATH)/src --proto_path=/usr/local/include ./*.proto

# }}} Protobuf end


# {{{ Cleanup
clean: protoclean

protoclean:
	rm -rf $(PROTO_GEN_FILES)
# }}} Cleanup end

# {{{ test

PROJECT_NAME := omilos-hub
PKG := gitlab.com/omilos/$(PROJECT_NAME)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/ | grep -v /models | grep -v /legacy)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)


test:
	go test $(PKG_LIST)

race:
	go test -race -short $(PKG_LIST)

coverage:
	coverage
	go tool cover -html=coverage.cov -o coverage.html

coverhtml:
	coverage
	go tool cover -html=coverage.cov -o coverage.html

lint: ## Lint the files
	echo ${PKG_LIST}
	go fmt ${PKG_LIST}
	go vet ${PKG_LIST}
	staticcheck ${PKG_LIST}

test-short:
	go test -short $(PKG_LIST)

# }}} test
