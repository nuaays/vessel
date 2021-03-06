### Makefile for tidb

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif

path_to_add := $(addsuffix /bin,$(subst :,/bin:,$(GOPATH)))
export PATH := $(path_to_add):$(PATH)

# Check the version of make and set env varirables/commands accordingly.
version_list := $(subst ., ,$(MAKE_VERSION))
major_version := $(firstword $(version_list))
old_versions := 0 1 2 3
ifeq "$(major_version)" "$(filter $(major_version),$(old_versions))"
  # Old version of `make` installed. It fails to search golex/goyacc
  # by using the modified `PATH`, so we specify these commands with full path.
  GODEP   = $$(which godep)
  GOLEX   = $$(which golex)
  GOYACC  = $$(which goyacc)
  GOLINT  = $$(which golint)
else
  # After version 4, `make` could follow modified `PATH` to find
  # golex/goyacc correctly.
  GODEP   := godep
  GOLEX   := golex
  GOYACC  := goyacc
  GOLINT  := golint
endif

GO        := $(GODEP) go
ARCH      := "`uname -s`"
LINUX     := "Linux"
MAC       := "Darwin"

LDFLAGS += -X "github.com/pingcap/tidb/util/printer.TiDBBuildTS=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "github.com/pingcap/tidb/util/printer.TiDBGitHash=$(shell git rev-parse HEAD)"

TARGET = ""

.PHONY: godep deps all build install parser clean todo test gotest interpreter server

all: godep parser build test check

godep:
	go get github.com/tools/godep
	go get github.com/pingcap/go-hbase
	go get github.com/pingcap/go-themis
	go get github.com/ngaut/tso/client

build:
	$(GO) build

install:
	$(GO) install ./...

TEMP_FILE = temp_parser_file

parser:
	go get github.com/qiuyesuifeng/goyacc
	go get github.com/qiuyesuifeng/golex
	$(GOYACC) -o /dev/null -xegen $(TEMP_FILE) parser/parser.y
	$(GOYACC) -o parser/parser.go -xe $(TEMP_FILE) parser/parser.y 2>&1 | egrep "(shift|reduce)/reduce" | awk '{print} END {if (NR > 0) {print "Find conflict in parser.y. Please check y.output for more information."; system("rm -f $(TEMP_FILE)"); exit 1;}}'
	rm -f $(TEMP_FILE)
	rm -f y.output

	@if [ $(ARCH) = $(LINUX) ]; \
	then \
		sed -i -e 's|//line.*||' -e 's/yyEofCode/yyEOFCode/' parser/parser.go; \
	elif [ $(ARCH) = $(MAC) ]; \
	then \
		/usr/bin/sed -i "" 's|//line.*||' parser/parser.go; \
		/usr/bin/sed -i "" 's/yyEofCode/yyEOFCode/' parser/parser.go; \
	fi

	$(GOLEX) -o parser/scanner.go parser/scanner.l

check:
	go get github.com/golang/lint/golint

	@echo "vet"
	@ go tool vet . 2>&1 | grep -vE 'Godeps|parser/scanner.*unreachable code' | awk '{print} END{if(NR>0) {exit 1}}'
	@echo "vet --shadow"
	@ go tool vet --shadow . 2>&1 | grep -vE 'Godeps' | awk '{print} END{if(NR>0) {exit 1}}'
	@echo "golint"
	@ $(GOLINT) ./... 2>&1 | grep -vE 'LastInsertId|NewLexer' | awk '{print} END{if(NR>0) {exit 1}}'
	@echo "gofmt (simplify)"
	@ gofmt -s -l . 2>&1 | grep -vE 'Godeps|parser/parser.go|parser/scanner.go' | awk '{print} END{if(NR>0) {exit 1}}'

deps:
	go list -f '{{range .Deps}}{{printf "%s\n" .}}{{end}}{{range .TestImports}}{{printf "%s\n" .}}{{end}}' ./... | \
		sort | uniq | grep -E '[^/]+\.[^/]+/' |grep -v "pingcap/tidb" | \
		awk 'BEGIN{ print "#!/bin/bash" }{ printf("go get -u %s\n", $$1) }' > deps.sh
	chmod +x deps.sh
	bash deps.sh

clean:
	$(GO) clean -i ./...
	rm -rf *.out
	rm -f deps.sh

todo:
	@grep -n ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alpha:]][[:alnum:]]* */*.go parser/scanner.l parser/parser.y || true
	@grep -n TODO */*.go parser/scanner.l parser/parser.y || true
	@grep -n BUG */*.go parser/scanner.l parser/parser.y || true
	@grep -n println */*.go parser/scanner.l parser/parser.y || true

test: gotest

gotest:
	$(GO) test -cover ./...

race:
	$(GO) test --race -cover ./...

interpreter:
	@cd interpreter && $(GO) build -ldflags '$(LDFLAGS)'

server:
ifeq ($(TARGET), "")
	@cd tidb-server && $(GO) build -ldflags '$(LDFLAGS)'
else
	@cd tidb-server && $(GO) build -ldflags '$(LDFLAGS)' -o '$(TARGET)'
endif
