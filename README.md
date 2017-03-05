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
You can install dependencies with the following `make` target:

```shell
make deps
```

Updating dependencies can be done as follows:

```shell
make deps-update
```

### Building

You can create a build by using the following `make` target:

```shell
# requires Golang with cross-compile support for Darwin and Linux
make build
```

This will create two builds in the `build` directory for Darwin and 
Linux targets. You can install them via the following `make` target:

```shell
# requires a writeable /usr/local/bin
make install
```

## Contributing

Pull requests are welcome! Herokuish is written in Bash and Go. Please conform to the [Bash styleguide](https://github.com/progrium/bashstyle) used for this project when writing Bash.

Developers should have Go installed with cross-compile support for Darwin and Linux. Tests will require Docker to be available. If you have OS X, we recommend boot2docker.

For help and discussion beyond Github Issues, join us on Freenode in `#gliderlabs`.

## Releases

Anybody can propose a release. First bump the version in `Makefile`, make sure `CHANGELOG.md` is up to date, and make sure tests are passing. Then open a Pull Request from `master` into the `release` branch. Once a maintainer approves and merges, CircleCI will build a release and upload it to Github.

### TODO

- Needs oh so much testing love
- Look into blackfriday's 2.0 API
