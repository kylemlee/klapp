FROM golang:latest as build
LABEL maintainer="Kyle Lee <kyleleemm@gmail.com>"
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR $GOPATH/src/github.com/kylemlee/klapp
COPY . .
RUN dep ensure && CGO_ENABLED=0 go build -a -o klapp
CMD ["./klapp"]
