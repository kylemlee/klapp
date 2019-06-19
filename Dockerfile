FROM golang:latest as build
LABEL maintainer="Kyle Lee <kyleleemm@gmail.com>"
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR $GOPATH/src/github.com/kylemlee/klapp
COPY . .
RUN dep ensure
CMD ["go","run","app.go"]