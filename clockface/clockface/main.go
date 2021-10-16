package main

import (
	"os"
	"time"

	"github.com/rjar2020/learn-go-with-tests/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
