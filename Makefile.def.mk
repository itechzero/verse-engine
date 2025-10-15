.ONESHELL:

SHELL := /bin/sh
PROJECT=$(shell basename ${PWD})
INTERFACE=$(shell netstat -r | head -n 3 | tail -n 1 | awk '{print $$NF}')
IP=$(shell ifconfig ${INTERFACE} | grep inet | grep -v inet6 | awk '{print $$2}')

OS=$(shell go env GOOS)
OS_ARCH=$(shell go env GOHOSTARCH)

GOLANGCI_VERSION=1.55.2

PROJECT_NAME = verse-toolkit

GIT_BRANCH = $(shell git branch | grep "*" | cut -d' ' -f2)
GIT_VERSION = $(shell git show -s --format='format:%h %aI')
BUILD ?= $(GIT_BRANCH) $(GIT_VERSION) $(shell go version)
GOBUILD = -ldflags '-X "$(PROJECT_CAPITAL_NAME)/pkg/version.Version=$(BUILD)"'
