CHS_ENV_HOME ?= $(HOME)/.chs_env
TESTS        ?= ./...

bin          := insolvency-poc
chs_envs     := $(CHS_ENV_HOME)/global_env $(CHS_ENV_HOME)/insolvency-poc/env
source_env   := for chs_env in $(chs_envs); do test -f $$chs_env && . $$chs_env; done
xunit_output := test.xml
lint_output  := lint.txt

commit       := $(shell git rev-parse --short HEAD)
tag          := $(shell git tag -l 'v*-rc*' --points-at HEAD)
version      := $(shell if [[ -n "$(tag)" ]]; then echo $(tag) | sed 's/^v//'; else echo $(commit); fi)

.EXPORT_ALL_VARIABLES:
GO111MODULE = on

.PHONY: all
all: build

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build: fmt
	go build

.PHONY: test
test: test-unit test-integration

.PHONY: test-unit
test-unit:
	go test $(TESTS) -run 'Unit' -coverprofile=coverage.out

.PHONY: test-integration
test-integration:
	$(source_env); go test $(TESTS) -run 'Integration'

.PHONY: clean
clean:
	go mod tidy
	rm -f ./$(bin) ./$(bin)-*.zip $(test_path) build.log

.PHONY: package
package:
	$(eval tmpdir := $(shell mktemp -d build-XXXXXXXXXX))
	cp ./$(bin) $(tmpdir)
	cp ./start.sh $(tmpdir)
	cd $(tmpdir) && zip -r ../$(bin)-$(version).zip $(bin) start.sh
	rm -rf $(tmpdir)

.PHONY: dist
dist: clean build package

.PHONY: xunit-tests
xunit-tests: GO111MODULE = off
xunit-tests:
	go get github.com/tebeka/go2xunit
	@set -a; go test -v $(TESTS) -run 'Unit' | go2xunit -output $(xunit_output)

.PHONY: lint
lint: GO111MODULE = off
lint:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	gometalinter ./... > $(lint_output); true