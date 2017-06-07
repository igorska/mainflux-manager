###
# Mainflux GTTP Sender Dockerfile
###

FROM golang
MAINTAINER Mainflux

ENV MONGO_HOST mongo
ENV MONGO_PORT 27017

ARG app_env
ENV APP_ENV $app_env

WORKDIR /go/src/github.com/mainflux/mainflux-manager

###
# Install
###
RUN go install

###
# Run main command with dockerize
###
# CMD mainflux-manager -m $MONGO_HOST
CMD if [ ${APP_ENV} = production ]; \
	then \
	mainflux-manager -m $MONGO_HOST; \
	else \
	go get github.com/pilu/fresh && \
	fresh -c mainflux-fresh.conf; \
	fi
