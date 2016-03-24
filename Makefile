deploy: vet lint fmt
	@git remote add deploy "${git_origin}"
	git push -u deploy -f master

vet:
	go tool vet $(shell pwd)

lint:
	golint

fmt:
	go fmt
	git diff --exit-code >/dev/null || echo "Go fmt found corrections"
	git diff --exit-code >/dev/null
