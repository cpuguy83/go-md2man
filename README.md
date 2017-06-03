# go-md2man

> ** Work in Progress **

This still needs a lot of help to be complete, or even usable!

Uses blackfriday to process markdown into man pages.

## Usage

By default, `md2man` takes input on stdin and sends the output to stdout:

```shell
cat README.md | md2man
```

You may optionally specify either a `-in` or `-out` parameters:

```shell
md2man -in README.md -out README.man
```

## Development

### Dependencies

We use [govend](https://github.com/govend/govend) for vendoring Go packages.
You can install or update dependencies with the following `make` target:

```shell
make deps
```

### Building

You can create a build by using the following `make` targets:

```shell
# builds a version for your OS (Darwin/Linux currently supported)
make build

# requires Golang with compile support for Darwin
# binary will be available in build/Darwin/md2man
make build-darwin

# requires Golang with compile support for Linux
# binary will be available in build/Linux/md2man
make build-linux

# requires Golang with cross-compile support for Darwin and Linux
# binaries will be available in build/Darwin/md2man and build/Linux/md2man
make build-cross
```

This will create two builds in the `build` directory for Darwin and 
Linux targets.

## Running tests

After installing dependencies, you can run tests by using the following
`make` target:

```shell
make test
```

## Contributing

Pull requests are welcome! go-md2man is written in Golang.

To ensure that builds work across all platforms, developers should have
Go installed with cross-compile support for Darwin and Linux.

## TODO

- Needs oh so much testing love
- Look into blackfriday's 2.0 API
