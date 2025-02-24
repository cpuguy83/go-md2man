[![GoDoc][go-docs-badge]][go-docs]
[![GoReportCard][go-report-card-badge]][go-report-card]
[![License][badge-license]][license]

go-md2man
=========

Converts markdown into roff (man pages).

Uses [blackfriday](https://github.com/russross/blackfriday) to process markdown into man pages.


### Installation

You can obtain the pre-built `go-md2man` binary for your OS and arch
[here](https://github.com/cpuguy83/go-md2man/releases); 

you can also install the `go-md2man` binary directly onto your `$GOBIN` dir with:
```sh
go install github.com/cpuguy83/go-md2man@latest
```

### Usage

```
go-md2man -in /path/to/markdownfile.md -out /manfile/output/path
```

See [go-md2man.1.md](./go-md2man.1.md) for an input example.

[license]: ./LICENSE.md
[badge-license]: https://img.shields.io/github/license/cpuguy83/go-md2man.svg
[go-docs-badge]: https://godoc.org/github.com/cpuguy83/go-md2man?status.svg
[go-docs]: https://godoc.org/github.com/cpuguy83/go-md2man/v2
[badge-build]: https://github.com/cpuguy83/go-md2man/actions/workflows/test.yml/badge.svg
[build]: https://github.com/cpuguy83/go-md2man/actions/workflows/test.yml
[go-report-card-badge]: https://goreportcard.com/badge/github.com/cpuguy83/go-md2man/v2
[go-report-card]: https://goreportcard.com/report/github.com/cpuguy83/go-md2man/v2
