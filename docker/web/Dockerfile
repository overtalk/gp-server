# build go program first
FROM golang as builder

WORKDIR $GOPATH/src/github.com/qinhan-shu/gp-server/

ADD . $GOPATH/src/github.com/qinhan-shu/gp-server

RUN export BRANCH="$(git rev-parse --abbrev-ref HEAD)" && \
    export COMMIT="$(git rev-parse --short HEAD)" && \
    export LDFLAGS=" -X main.commit=$COMMIT -X main.branch=$BRANCH" && \
    GO15VENDOREXPERIMENT=1 CGO_ENABLED=0 go build -ldflags "$LDFLAGS" -o ./build/gp-server ./server/web/server.go

RUN cp ./build/gp-server /usr/local/bin/

# build php container
FROM shiyunjin/docker-supervisor:latest

ENV GITHUB_URL=""
        
# Put other config and shell files into place.
COPY docker/web/supervisord.conf /etc/supervisor/supervisord.conf

# Put go executeable to bin
COPY --from=builder /usr/local/bin/gp-server /usr/local/bin/gp-server 

EXPOSE 8080
