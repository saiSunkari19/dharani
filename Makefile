PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=dharani \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=dharanid \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=dharanicli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

install: go.sum
		go install  $(BUILD_FLAGS) ./cmd/dharanid
		go install  $(BUILD_FLAGS) ./cmd/dharanicli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

init-pre:
	rm -rf ~/.dharanicli
	rm -rf ~/.dharanid
	dharanid init mynode --chain-id dharani
	dharanicli config keyring-backend test

init-user1:
	dharanicli keys add user1 --output json 2>&1

init-user2:
	dharanicli keys add user2 --output json 2>&1
