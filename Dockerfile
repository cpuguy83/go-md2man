FROM golang:1.8 AS build
ENV NAME=go-md2man \
    BINARYNAME=md2man \
    OWNER=cpuguy83 \
    CGO_ENABLED=0

COPY . "/go/src/github.com/$OWNER/$NAME"
WORKDIR "/go/src/github.com/$OWNER/$NAME"
RUN VERSION="$(git rev-parse --abbrev-ref HEAD)" \
    && go build -ldflags "-X main.Version=$VERSION"

FROM scratch
COPY --from=build /go/src/github.com/cpuguy83/go-md2man/go-md2man /bin/md2man
ENTRYPOINT ["/bin/md2man"]
