BINARY := web-teller
BUILDCMD := go build -trimpath -ldflags="-s -w" -o
OUTPUTDIR := build
OUTPUT := ./$(OUTPUTDIR)/$(BINARY)
tag := 1.0

.PHONY: build
build:
	GIT_TERMINAL_PROMPT=1 GO111MODULE=on GOPRIVATE=gitlab.pactindo.com CGO_ENABLED=0 $(BUILDCMD) $(OUTPUT) main.go && upx $(OUTPUT)
