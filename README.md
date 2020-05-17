# Rich Contextual Go Errors

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/awfm/rich) [![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/awfm/rich/master/LICENSE) [![Build Status](https://travis-ci.org/awfm/rich.svg?branch=master)](https://travis-ci.org/awfm/rich)

The rich package provides errors with structured context information for popular Go logging libraries:

- [Zerolog](https://github.com/rs/zerolog)
- [Zap](https://github.com/uber-go/zap)
- [Logrus](https://github.com/sirupsen/logrus)

## Installation

```sh
go get -u github.com/awfm/rich
```

## Examples

### Zerolog

```go
package main

import (
  "io"
  "os"

  rich "github.com/awfm/rich/zerolog"
  "github.com/rs/zerolog/log"
)

func main() {

  log := zerolog.New(os.Stderr)

  var src, dst *os.File

  err := copyFile(src, dst)
  if err != nil {
    rich.Log(log.Fatal).
      Err(err).
      Str("src", src.Name()).
      Str("dst", dst.Name()).
      Msg("could not copy file")
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

### Zap

```go
package main

import (
  "io"
  "os"

  rich "github.com/awfm/rich/zap"
  "go.uber.org/zap"
)

func main() {

  log := zap.NewProduction()

  var src, dst *os.File

  err := copyFile(src, dst)
  if err != nil {
    rich.Log(log).
      With(
        zap.Error(err), // has to be first item to use rich error context
        zap.String("src", src.Name()),
        zap.String("dst", src.Name()),
      ).
      Fatal("could not copy file")
  }

  os.Exit(0)
}

func copyFile(src *os.File, dst *os.File) error {

  n, err := io.Copy(src, dst)
  if err != nil {
    return rich.Errorf("could not copy contents: %w", err).With(zap.Int64("bytes_written", n))
  }
  
  return nil
}
```

### Zap (sugared)

```go
package main

import (
  "io"
  "os"

  rich "github.com/awfm/rich/zap"
  "go.uber.org/zap"
)

func main() {

  sugar := zap.NewProduction().Sugar()

  var src, dst *os.File

  err := copyFile(src, dst)
  if err != nil {
    rich.Sugar(sugar).
      With(
        "error", err, // has to be first item to use rich error context
        "src", src.Name(),
        "dst", dst.Name(),
      ).
      Fatal("could not copy file")
  }

  os.Exit(0)
}

func copyFile(src *os.File, dst *os.File) error {

  n, err := io.Copy(src, dst)
  if err != nil {
    return rich.Errorf("could not copy contents: %w", err).Sugar().With("bytes_written", n)
  }
  
  return nil
}
```

### Logrus

```go
package main

import (
  "io"
  "os"

  rich "github.com/awfm/rich/logrus"
  "github.com/sirupsen/logrus"
)

func main() {

  log := logrus.New()

  var src, dst *os.File

  err := copyFile(src, dst)
  if err != nil {
    rich.Log(log).
      WithError(err).
      WithFields(logrus.Fields{
        "src": src.Name(),
        "dst", dst.Name(),
      }).
      Fatal("could not copy file")
  }

  os.Exit(0)
}

func copyFile(src *os.File, dst *os.File) error {

  n, err := io.Copy(src, dst)
  if err != nil {
    return rich.Errorf("could not copy contents: %w", err).WithField("bytes_written", n)
  }
  
  return nil
}
```

### Output

```json
{"level": "fatal", "src": "file1", "dst": "file2", "bytes_written": 123, "err": "could not copy contents: some file error"}
```

## Rationale

I like Go error handling. While verbose, it offers a pragmatic and unambiguous way to handle failure. With the introduction of error wrapping into the standard library in Go 1.13, we now even have a portable way to pass contextual information for errors across API boundaries.

I also love structured JSON logging. With tools such as `jq`, it becomes easy to analyze your application's logic in detail. The visibility gained is an indispensable part of monitoring and maintenance in large-scale deployments of production applications.

Unfortunately, while both the logging libraries and the error wrapping idiom offer ways to add context to error messages, they are not compatible in a meaningful way. The rich library was created to bridge this gap and log the contextual information of an error in a structured manner.

For instance, consider some of the code from the Zerolog example above, when not using `rich`:

```go
n, err := io.Copy(in, out)
if err != nil {
  return fmt.Errorf("could not copy contents (bytes written: %d): %w", n, err)
}
```

When logging the error, the context - unknown to the calling function - cannot be part of structured logging:

```go
if err != nil {
  log.Error().Str("src", src).Str("dst", dst).Err(err).Msg("could not copy file")
}
```

The output will mix structured contextual information from the caller with unstructured error context:

```json
{"level": "fatal", "src": "file1", "dst": "file2", "err": "could not copy contents (bytes written: 123): some file error"}
```

The rich error package bridges this gap between error context and logging context, bringing all of the advantages of structured logging to errors.

## Tips

Error handling with go is simple; however, there are still a few tips to keep in mind to get the most out of it.

### Don't add function parameters to the error context

When a function is called, the caller already has access to all the information on the parameters. It should therefore be left to the caller which information is included in the context for logging.

Don't do:

```go
func do(p1 string, p2 uint64) error {
  return rich.Errorf("could not do stuff: %w", err).Str("p1", p1).Uint64("p2", p2)
}
```

Instead, do:

```go
func do(p1 string, p2 uint64) error {
   return rich.Errorf("could not do stuff: %w", err)
}
```

The caller can than choose:

```go
err := do(p1, p2)
if err != nil {
  rich.Log(log.Warn).Str("p1", p1).Uint64("p2", p2)
}
```

### Only provide context relevant for the error path

When you have multiple error paths in a function, don't include all of the information in each of them. If context information isn't relevant for a path, don't include it.

If you have this:

```go
n, err := f.Write(data)
if err != nil {
  return rich.Errorf("could not write data: %w", err).Int64("bytes_written", n)
}
```

Don't do:

```go
err = f.Close()
if err != nil {
  return rich.Errorf("could not close file: %w", err).Int64("bytes_written", n)
}
```

Instead, do:

```go
err = f.Close()
if err != nil {
  return rich.Errorf("could not close file: %w", err)
}
```
