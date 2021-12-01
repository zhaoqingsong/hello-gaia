PACKAGES = \
github.com/zhaoqingsong/studyweb
RELEASE_PLATFORMS = \
darwin/amd64 \
linux/amd64
DOCKER_IMAGE = harbor.hellobike.cn/kube-public/appcontainerupladerservice
GOPROXY = http://athens.hellobike.cn:3000

VERSION := $(shell git describe --tags 2>/dev/null || (printf "v0.0.0-" && git rev-parse --short HEAD))
PROJECTNAME := $(shell basename "$(PWD)")
PROJECTNAME_LOWER := $(shell basename "$(PWD)" | tr '[:upper:]' '[:lower:]')
DEPLOY_NAMESPACE := "defensor"
DEPLOY_ENV := "dev"

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=main.Version=$(VERSION)"

PLATFORMS = \
darwin/386 \
darwin/amd64 \
darwin/arm \
darwin/arm64 \
dragonfly/amd64 \
freebsd/386 \
freebsd/amd64 \
freebsd/arm \
linux/386 \
linux/amd64 \
linux/arm \
linux/arm64 \
linux/ppc64 \
linux/ppc64le \
linux/mips \
linux/mipsle \
linux/mips64 \
linux/mips64le \
netbsd/386 \
netbsd/amd64 \
netbsd/arm \
openbsd/386 \
openbsd/amd64 \
openbsd/arm \
plan9/386 \
plan9/amd64 \
solaris/amd64 \
windows/386 \
windows/amd64

## build: Build for current environment.
build:
	@for pkg in $(PACKAGES); do \
		env GOPROXY=$(GOPROXY) go build $(LDFLAGS) $${pkg}; \
	done

## install: Install to the current environment.
install:
	@for pkg in $(PACKAGES); do \
		env GOPROXY=$(GOPROXY) go install $(LDFLAGS) $${pkg}; \
	done

## darwin/386: Build for GOOS=darwin GOARCH=386
## darwin/amd64: Build for GOOS=darwin GOARCH=amd64
## darwin/arm: Build for GOOS=darwin GOARCH=arm
## darwin/arm64: Build for GOOS=darwin GOARCH=arm64
## dragonfly/amd64: Build for GOOS=dragonfly GOARCH=amd64
## freebsd/386: Build for GOOS=freebsd GOARCH=386
## freebsd/amd64: Build for GOOS=freebsd GOARCH=amd64
## freebsd/arm: Build for GOOS=freebsd GOARCH=arm
## linux/386: Build for GOOS=linux GOARCH=386
## linux/amd64: Build for GOOS=linux GOARCH=amd64
## linux/arm: Build for GOOS=linux GOARCH=arm
## linux/arm64: Build for GOOS=linux GOARCH=arm64
## linux/ppc64: Build for GOOS=linux GOARCH=ppc64
## linux/ppc64le: Build for GOOS=linux GOARCH=ppc64le
## linux/mips: Build for GOOS=linux GOARCH=mips
## linux/mipsle: Build for GOOS=linux GOARCH=mipsle
## linux/mips64: Build for GOOS=linux GOARCH=mips64
## linux/mips64le: Build for GOOS=linux GOARCH=mips64le
## netbsd/386: Build for GOOS=netbsd GOARCH=386
## netbsd/amd64: Build for GOOS=netbsd GOARCH=amd64
## netbsd/arm: Build for GOOS=netbsd GOARCH=arm
## openbsd/386: Build for GOOS=openbsd GOARCH=386
## openbsd/amd64: Build for GOOS=openbsd GOARCH=amd64
## openbsd/arm: Build for GOOS=openbsd GOARCH=arm
## plan9/386: Build for GOOS=plan9 GOARCH=386
## plan9/amd64: Build for GOOS=plan9 GOARCH=amd64
## solaris/amd64: Build for GOOS=solaris GOARCH=amd64
## windows/386: Build for GOOS=windows GOARCH=386
## windows/amd64: Build for GOOS=windows GOARCH=amd64
$(PLATFORMS):
	@mkdir -p release
	@$(eval GOOS := $(shell echo $@ | awk -F/ '{print $$1}'))
	@$(eval GOARCH := $(shell echo $@ | awk -F/ '{print $$2}'))

	$(info Building for $(GOOS) $(GOARCH)...)

	@$(eval GENERATENAME := -$(GOOS)_$(GOARCH)-$(VERSION))
	@$(eval EXT := $(shell [ $(GOOS) = 'windows' ] && echo '.exe' || echo ''))
	@$(eval OUTPUTNAME := $(GENERATENAME)$(EXT))

	@for pkg in $(PACKAGES); do \
		PACKAGENAME=$${pkg##*/}; \
		env GOOS=$(GOOS) GOARCH=$(GOARCH) GOPROXY=$(GOPROXY) go build $(LDFLAGS) -o release/$${PACKAGENAME}$(OUTPUTNAME) $${pkg}; \
	done

## release: Make release for target platforms.
release: $(RELEASE_PLATFORMS)
	@echo Done.

## docker-image: Make docker image.
docker-image:
	@docker build --rm -t $(DOCKER_IMAGE):$(VERSION) .

## docker-push: Push the docker image to a remote registry
docker-push:
	docker push $(DOCKER_IMAGE):$(VERSION)

## deploy: Deploy the chart to the dest kubernetes cluster
deploy:
	helm --kubeconfig ~/.kube/config-$(DEPLOY_ENV) upgrade --install --namespace $(DEPLOY_NAMESPACE) $(PROJECTNAME_LOWER) -f ./chart/values-$(DEPLOY_ENV).yaml --set "annotations.revision=release-$$(date +%Y%m%d%H%M%S),image.tag=$(VERSION)" --set-file "config=./config_$(DEPLOY_ENV).yaml" ./chart/$(PROJECTNAME_LOWER)

## clean: Clean build files.
clean:
	@rm -rf release
	@go clean

## check: Check code format, vet, lint code etc.
check:
	go mod tidy
	go fmt ./...
	go vet
	golint ./...

## test: Test does all pre-release works.
test:
	go test -v ./...

.PHONY: help
all: help
## help: List all supported make commands.
help: Makefile
	@echo
	@echo "Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
