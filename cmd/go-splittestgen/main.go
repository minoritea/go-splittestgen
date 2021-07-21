package main

import (
	"flag"
	"fmt"
	"github.com/alessio/shellescape"
	"github.com/minoritea/go-splittestgen"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	total := flag.Uint("total", 1, "total number of test processes")
	index := flag.Uint("index", 0, "index of test processes (default 0, must be less than the total number)")
	flag.Parse()

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	tests := splittestgen.
		GetPackages(string(input)).
		Tests().
		DevideEquallyBy(int(*total))[*index]

	var commandStrs []string
	for _, cmd := range tests.Commands() {
		commandStrs = append(commandStrs,
			shellescape.QuoteCommand(
				append(
					[]string{
						"go",
						"test",
					},
					cmd.Args()...,
				),
			),
		)
	}

	_, err = fmt.Println(
		strings.Join(commandStrs, "\n"),
	)
	if err != nil {
		log.Fatal(err)
	}
}
