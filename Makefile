.PHONY: vet lint fmt test deploy

build: fmt vet lint test go-domainglass

go-domainglass: *.go
	@echo "== Building"
	@go build

deploy:
	@git remote add deploy "${git_origin}"
	git push -u deploy -f master

vet:
	@echo "== Vetting"
	@go tool vet $(shell pwd)

lint:
	@echo "== Linting"
	@golint

fmt:
	@echo "== Fmting"
	@if [ -n "$(shell gofmt -d -s .)" ]; then \
		gofmt -d -s .; \
		exit 1; \
	fi

test:
	@ echo "== Testing"
	@go test -covermode=count

coveralls:
	@goveralls -v -repotoken ${coveralls_token}
