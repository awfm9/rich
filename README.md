# Rich Contextual Errors for Zerolog

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/awfm/rich) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/awfm/rich/master/LICENSE) [![Build Status](https://travis-ci.org/awfm/rich.svg?branch=master)](https://travis-ci.org/awfm/rich) [![Coverage](http://gocover.io/_badge/github.com/awfm/rich)](http://gocover.io/github.com/awfm/rich)

The rich package provides errors with rich, structured context information for Zerolog.

## Advantages

- Useable as drop-in replacement to log normal errors
- Allows adding structured context fields to errors
- Can bubble up error context through API boundaries
- Conserves support for Go 1.13 error comparisons

## Installation

```sh
go get -u github.com/awfm/rich
```

## Example

```go
package main

import (
  "io"
  "os"

  "github.com/awfm/rich"
  "github.com/rs/zerolog"
)

func main() {

  log := zerolog.New(os.Stderr)

  var src, dst *os.File

  err := copyFile(src, dst)
  if err != nil {
    rich.Log(log.Fatal).Err(err).Str("src", src.Name()).Str("dst", dst.Name()).Msg("could not copy file")
  }

  os.Exit(0)
}

func copyFile(src *os.File, dst *os.File) error {

  n, err := io.Copy(src, dst)
  if err != nil {
    return rich.Errorf("could not copy contents: %w", err).Int64("bytes_written", n)
  }
  
  return nil
}
```

Output:

```json
{"level": "fatal", "src": "file1", "dst": "file2", "bytes_written": 123, "err": "could not copy contents: some file error"}
```

## Explanation

I like Go error handling. While verbose, it's a pragmatic and unambiguous way to handle failure. With the introduction of error wrapping into the standard library in Go 1.13, we even have a standard way to add information to errors in a human-readable manner.

I also love structured JSON logging. With tools such as `jq`, it becomes easy to analyze your applications logic in detail. The visibility gained is an indispensable part of monitoring and maintenance for large-scale production applications.

The issue is that, in the case of error paths, both the logging and the error wrapping will often be used to add contextual information. In an idea world, we would be able to bring them both together - and this is what `rich.Error` and `rich.Log` do.

Consider some of the code from the example without the use of the rich package:

```go
n, err := io.Copy(in, out)
if err != nil {
  return fmt.Errorf("could not copy contents (bytes written: %d): %w", n, err)
}
```

Now, when logging the error, the context that is invisible to the calling function can not be part of structured logging:

```go
if err != nil {
  log.Error().Str("src", src).Str("dst", dst).Err(err).Msg("could not copy file")
}
```

The output in this case would be mixing structured context with unstructured context embedded in the error:

```json
{"level": "fatal", "src": "file1", "dst": "file2", "err": "could not copy contents (bytes written: 123): some file error"}
```

The rich error package bridges this gap between error context and logging context.
