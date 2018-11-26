.PHONY: all

F3_API_CI_DOCKERFILE = ops/docker/ci.Dockerfile
F3_API_CI_IMAGE = hmoragrega/f3-payments-ci
F3_API_CI_TAG := 1.0.1

ci-build:
	@docker build -t ${F3_API_CI_IMAGE}:${F3_API_CI_TAG} -t ${F3_API_CI_IMAGE}:latest -f ${F3_API_CI_DOCKERFILE} .

ci-push:
	@docker push ${F3_API_CI_IMAGE}:${F3_API_CI_TAG}
	@docker push ${F3_API_CI_IMAGE}:latest

ci-run:
	@docker run -it --entrypoint /bin/sh ${F3_API_CI_IMAGE}:${F3_API_CI_TAG}