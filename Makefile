default: build

VERSION ?= 5.1

GITSHA ?= "no-git-module"
ifneq ("$(wildcard .git)", "")
	GITSHA = $(shell git rev-parse --short=7 HEAD)
endif

VERSYM="github.com/tao12345666333/argo-cd-demo/pkg/version._buildVersion"
GITSHASYM="github.com/tao12345666333/argo-cd-demo/pkg/version._buildGitRevision"
GO_LDFLAGS ?= "-X=$(VERSYM)=$(VERSION) -X=$(GITSHASYM)=$(GITSHA)"

.PHONY: build
build:
	CGO_ENABLED=0 go build \
		-o argo-cd-demo \
		-ldflags $(GO_LDFLAGS) \
		main.go

