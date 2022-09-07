GO ?= go

.PHONY: debug
debug:
	 $(GO) build -o ${GOPATH}/bin/gdev-i18n gdev-i18n.go
	 $(GO) generate ./...