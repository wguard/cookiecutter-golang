package main

import (
	"flag"
	"fmt"
	"os"
)

func run() error {
	return nil
}

var version string
var commit string
var build string

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "print version")
	flag.Parse()

	if showVersion {
		fmt.Fprintf(os.Stderr, "%s(commit:%s build at:%s)", version, commit, build)
		os.Exit(0)
	}

	if err := run(); err != nil {
		panic(err)
	}
}
