BINARY := web-teller
BUILDCMD := go build -trimpath -ldflags="-s -w" -o
OUTPUTDIR := build
OUTPUT := ./$(OUTPUTDIR)/$(BINARY)
tag := 1.1
TIME := $(shell date +%Y%m%d)
PATHNEXUS := ebanking/ibmb

.PHONY: build
build:
	GIT_TERMINAL_PROMPT=1 GO111MODULE=on GOPRIVATE=gitlab.pactindo.com CGO_ENABLED=0 $(BUILDCMD) $(OUTPUT) main.go && upx $(OUTPUT)

.PHONY: linux
linux:
	GOOS=linux GOARCH=amd64 $(BUILDCMD) $(OUTPUT) main.go && upx $(OUTPUT)

.PHONY: docker
docker:
	docker build -t $(BINARY):${tag}.$(TIME) .
	docker tag $(BINARY):${tag}.$(TIME) nexus.pactindo.com:8443/${PATHNEXUS}/$(BINARY):${tag}.$(TIME)
	docker push nexus.pactindo.com:8443/${PATHNEXUS}/$(BINARY):${tag}.$(TIME)

