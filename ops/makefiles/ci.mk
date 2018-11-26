.PHONY: all

CI_DOCKERFILE = ops/docker/ci.Dockerfile
CI_IMAGE = hmoragrega/f3-payments-ci
CI_TAG := 1.0.3

ci-build:
	@docker build -t ${CI_IMAGE}:${CI_TAG} -t ${CI_IMAGE}:latest -f ${CI_DOCKERFILE} .

ci-push:
	@docker push ${CI_IMAGE}:${CI_TAG}
	@docker push ${CI_IMAGE}:latest

ci-run:
	@docker run -it --entrypoint /bin/sh ${CI_IMAGE}:${CI_TAG}