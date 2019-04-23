FROM golang:latest

WORKDIR $GOPATH/src/github.com/qinhan-shu/gp-server/

ADD ./. .

ENV GOOS=linux
ENV GO15VENDOREXPERIMENT=1
ENV CGO_ENABLED=0

RUN export BRANCH="$(git rev-parse --abbrev-ref HEAD)" && \
    export COMMIT="$(git rev-parse --short HEAD)" && \
    export LDFLAGS=" -X main.commit=$COMMIT -X main.branch=$BRANCH" && \
    go build -ldflags "$LDFLAGS" -o ./build/gp-server ./server/web/server.go

CMD ./build/gp-server  -uploadPath $GOPATH/src/github.com/qinhan-shu/gp-server/docker_data/judge/judgefile/ 

EXPOSE 8080
