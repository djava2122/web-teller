BINARY := web-teller
BUILDCMD := go build -trimpath -ldflags="-s -w" -o
OUTPUTDIR := build
OUTPUT := ./$(OUTPUTDIR)/$(BINARY)
TIME := $(shell date +%Y%m%d)
PATHNEXUS := ebanking/ibmb

.PHONY: build
build:
	GIT_TERMINAL_PROMPT=1 GO111MODULE=on GOPRIVATE=git.pactindo.com CGO_ENABLED=0 $(BUILDCMD) $(OUTPUT) main.go && upx $(OUTPUT)

.PHONY: linux
linux:
	GOOS=linux GOARCH=amd64 $(BUILDCMD) $(OUTPUT) main.go && upx $(OUTPUT)

.PHONY: docker
docker:
	docker build -t $(BINARY):${tag} .
	docker tag $(BINARY):${tag} nexus.pactindo.com:8443/${PATHNEXUS}/$(BINARY):${tag}
	docker push nexus.pactindo.com:8443/${PATHNEXUS}/$(BINARY):${tag}

