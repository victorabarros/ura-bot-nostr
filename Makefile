APP_NAME?=$(shell pwd | xargs basename)
APP_DIR = /api/${APP_NAME}
PWD=$(shell pwd)
IMAGE?=golang:1.22.3
COMMAND?=bash

YELLOW=$(shell printf '\033[0;1;33m')
COLOR_OFF=$(shell printf '\033[0;1;0m')

debug:
	@echo "${YELLOW}Debug Mode${COLOR_OFF}"
	@docker run -it \
		--env-file .env \
		-v ${PWD}:${APP_DIR} \
		-w ${APP_DIR} \
		-p 8092:3000 \
		--rm --name ${APP_NAME} ${IMAGE} ${COMMAND}
