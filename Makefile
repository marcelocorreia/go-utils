test: consul-dev-start go-test
	$(MAKE) consul-dev-stop

go-test:
	go test $$( glide nv)

deps:
	glide install

deps-update:
	glide update
