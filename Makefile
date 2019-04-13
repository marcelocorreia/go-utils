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

snapshot:
	-mkdir -p dist coverage
	goreleaser  release --snapshot  --rm-dist --debug

release:
	goreleaser release

_dep-ensure:
	dep ensure

_open-page:
	open https://github.com/$(GITHUB_USER)/$(GIT_REPO_NAME).git

_open-coverage:
	open ./coverage/index.html