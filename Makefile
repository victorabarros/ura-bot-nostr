APP_NAME?=$(shell pwd | xargs basename)
APP_DIR = /go/src/github.com/victorabarros/${APP_NAME}
PWD=$(shell pwd)
DOCKER_BASE_IMAGE=golang:1.22
PORT=8092
COMMAND?=bash

YELLOW=$(shell printf '\033[0;1;33m')
COLOR_OFF=$(shell printf '\033[0;1;0m')

docker-cmd:
	@echo "${YELLOW}Debug Mode${COLOR_OFF}"
	@docker run -it \
		--env-file .env \
		-v ${PWD}:${APP_DIR} \
		-w ${APP_DIR} \
		-p ${PORT}:${PORT} \
		--rm --name ${APP_NAME}-docker-cmd \
		${DOCKER_BASE_IMAGE} bash -c "${COMMAND}"

test:
	@clear
	@make docker-cmd COMMAND="go test -v -coverprofile=c.out"

coverage:
	@go tool cover -html=c.out
