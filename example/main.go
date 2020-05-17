package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/awishformore/rich"
	"github.com/rs/zerolog"
)

func main() {

	log := zerolog.New(os.Stderr)

	p1 := rand.Float32()
	p2 := rand.Uint64()

	err := do(p1, p2)
	if err != nil {
		rich.Log(log.Warn).Err(err).Float32("p1", p1).Uint64("p2", p2).Msg("could not do stuff")
	}

	fmt.Println(err)

	fmt.Println(errors.Is(err, io.EOF))
}

func do(p1 float32, p2 uint64) error {

	timeout := 1*time.Minute + 23*time.Second + 456*time.Millisecond
	timestamp := time.Now().UTC()

	return rich.Errorf("could not read message: %w", io.EOF).Dur("timeout", timeout).Time("timestamp", timestamp)
}
