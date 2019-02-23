# docker builder image
FROM golang:alpine as builder

COPY . /go/src/github.com/secanis/docker-image-checker
WORKDIR /go/src/github.com/secanis/docker-image-checker
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o dic .

# docker publisher image
FROM golang:latest as publisher

ARG GITHUB_TOKEN

WORKDIR /
RUN mkdir -p /goreleaser \
    && curl --silent "https://api.github.com/repos/goreleaser/goreleaser/releases/latest" | \
    grep '"tag_name":' | \
    sed -E 's/.*"([^"]+)".*/\1/' | \
    xargs -I {} curl -o /goreleaser_Linux_x86_64.tar.gz -sOL "https://github.com/goreleaser/goreleaser/releases/download/"{}'/goreleaser_Linux_x86_64.tar.gz' \
    && tar -xzf goreleaser_Linux_x86_64.tar.gz -C /goreleaser
COPY . /go/src/github.com/secanis/docker-image-checker
WORKDIR /go/src/github.com/secanis/docker-image-checker
RUN /goreleaser/goreleaser release


# final docker image
FROM alpine

# add user and install certs
RUN adduser -S -D -H -h /app appuser \
    && apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
USER appuser
COPY --from=builder /go/src/github.com/secanis/docker-image-checker/dic /app/
WORKDIR /app

CMD ["./dic", "-h"]