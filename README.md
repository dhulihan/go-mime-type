# go-mime-type

A simple CLI tool that prints mime type, as recognized by [http.DetectContentType](https://pkg.go.dev/net/http#DetectContentType).

## Install

```sh
go install github.com/dhulihan/go-mime-type
```

## Usage

```sh
go-mime-type <path>
```

Example:

```sh
$ go-mime-type my-file.opus
application/ogg
```

## Why?

Go's `http.DetectContentType` often disagrees with other mime type detection tools (eg: `file --mime-type`), so this tool helps you identify those disagreements.
