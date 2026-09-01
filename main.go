package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gregyr/go-utils/cli"
	"github.com/gregyr/go-utils/std"
	"github.com/gregyr/go-utils/streams"
)

const ()

func main() {
	c := cli.New("testapp", "file")

	c.AddFlag(cli.NewFlag("v", "verbose", "prints the full output"))
	c.AddFlag(cli.NewFlag("e", "even", "only returns even numbers"))
	c.AddOpt(cli.NewOpt("o", "output", "outputs the result into the provided file instead of stdout"))

	err := c.LoadArgs()
	if err != nil {
		panic(err)
	}

	run(c)
}

func run(c cli.CLI) {

	in := c.GetInput()
	n := 10
	if len(in) > 0 {
		number, err := strconv.Atoi(in)
		if err != nil {
			panic("Unable to parse input")
		}
		n = number
	}

	r := std.Range(n)
	if c.Has("verbose") {
		fmt.Printf("created: %v\n", r)
		fmt.Println("Processing...")
	}

	if c.Has("even") {
		if c.Has("v") {
			fmt.Println("Streaming...")
		}
		r = streams.StreamSlice(r).Filter(func(a int) bool { return a%2 == 0 }).ToSlice()
		if c.Has("v") {
			fmt.Println("Filtered to evens")
			fmt.Println(r)
		}
	}

	output := c.GetOption("o")
	if output != "" {
		f, err := os.Create(output)
		if err != nil {
			panic(err)
		}
		for _, n := range r {
			fmt.Fprintln(f, n)
		}
	} else {
		for _, n := range r {
			fmt.Println(n)
		}
	}
}
