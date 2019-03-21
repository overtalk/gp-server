VERSION := $(shell git describe --exact-match --tags 2>/dev/null)

# 可执行文件的位置
PREFIX := /usr/local
# 获取git分支
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
# 获取git提交版本
COMMIT := $(shell git rev-parse --short HEAD)


ifdef GOBIN
PATH := $(GOBIN):$(PATH)
else
PATH := $(subst :,/bin:,$(shell go env GOPATH))/bin:$(PATH)
endif

LDFLAGS := $(LDFLAGS) -X main.commit=$(COMMIT) -X main.branch=$(BRANCH)
ifdef VERSION
	LDFLAGS += -X main.tag=$(VERSION)
endif

.PHONY: deps
deps:
	go get ./...

.PHONY: server
server:
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o build/gp_server ./server/web/server.go

.PHONY: install
install: server
	mkdir -p $(DESTDIR)$(PREFIX)/bin/
	cp build/gp_server $(DESTDIR)$(PREFIX)/bin/

.PHONY: clean
clean:
	rm -f build/gp_server
