package main

import (
	"errors"
	"flag"
	"fmt"
)

type Foo struct {
	bananaPrefix string
	bananaScheme string
	ignoreApple  bool
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	var prefix = flag.String("banana-prefix", "", "prefix of the banana")
	var scheme = flag.String("banana-scheme", "", "banana scheme to use")
	var ignoreApple = flag.Bool("ignore-apple", false, "ignore all apples")

	flag.Parse()

	return validate(&Foo{
		bananaPrefix: *prefix,
		bananaScheme: *scheme,
		ignoreApple:  *ignoreApple,
	})
}

func validate(f *Foo) error {
	if len(f.bananaPrefix) == 0 || len(f.bananaScheme) == 0 || !f.ignoreApple {
		return errors.New("invalid Foo")
	}
	return nil
}

