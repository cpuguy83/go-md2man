FROM golang:1.8 AS build
COPY . /go/src/github.com/cpuguy83/go-md2man
WORKDIR /go/src/github.com/cpuguy83/go-md2man
RUN VERSION="$(git rev-parse --abbrev-ref HEAD)" \
    && CGO_ENABLED=0 go build -ldflags "-X main.Version=$VERSION"

FROM scratch
COPY --from=build /go/src/github.com/cpuguy83/go-md2man/go-md2man /go-md2man
ENTRYPOINT ["/go-md2man"]
