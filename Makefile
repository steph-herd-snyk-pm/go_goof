.PHONY: FORCE
FORCE: # targets depending on FORCE will always be built (see https://www.gnu.org/software/make/manual/html_node/Force-Targets.html)

snyk-test: build FORCE
	docker run -e SNYK_TOKEN=${SNYK_TOKEN} -v $(shell pwd):/src -w /src snyk-test-go-dummy

go-version: build FORCE
	docker run snyk-test-go-dummy bash -c "go version"

manifest: build FORCE
	docker run -v $(shell pwd):/src -w /src snyk-test-go-dummy bash -c "go list -json -deps > gomod.manifest"

build: FORCE
	docker build -t snyk-test-go-dummy .
