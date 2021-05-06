#test: consul-dev-start go-test
#	$(MAKE) consul-dev-stop
#
#go-test:
#	go test $$( glide nv)
#
#deps:
#	glide install
#
#deps-update:
#	glide update

SCAF := go run cmd/scafold/main.go
GITHUB_USER ?= marcelocorreia
GIT_REPO_NAME ?= go-utils
SEMVER_DOCKER ?= marcelocorreia/semver
RELEASE_TYPE ?= patch


wrap-up:
	go mod tidy
	#go mod vendor


snapshot:
	-mkdir -p dist coverage
	goreleaser  release --snapshot  --rm-dist --debug

release: _setup-versions _tag-push
	goreleaser release  --rm-dist

_tag-push:
	-git add .
	-git commit -m "Release: $(NEXT_VERSION)"
	-git tag $(NEXT_VERSION)
	-git push
	-git push --tags

all-versions:
	@git ls-remote --tags $(GIT_REMOTE)

current-version: _setup-versions
	@echo $(CURRENT_VERSION)

next-version: _setup-versions
	@echo $(NEXT_VERSION)

_setup-versions:
	$(eval export CURRENT_VERSION=$(shell git ls-remote --tags $(GIT_REMOTE) | grep -v latest | awk '{ print $$2}'|grep -v 'stable'| sort -r --version-sort | head -n1|sed 's/refs\/tags\///g'))
	$(eval export NEXT_VERSION=$(shell docker run --rm --entrypoint=semver $(SEMVER_DOCKER) -c -i $(RELEASE_TYPE) $(CURRENT_VERSION)))

_dep-ensure:
	dep ensure

_open-page:
	open https://github.com/$(GITHUB_USER)/$(GIT_REPO_NAME).git

_open-coverage:
	open ./coverage/index.html

define git_push
	-git add .
	-git commit -m "$1"
	-git push
endef