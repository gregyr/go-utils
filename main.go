package main

import (
	"fmt"

	cmd "github.com/gregyr/go-utils/framework/cli/full"
)

type c1Args struct {
	oye     *bool
	someval *int
}

func main() {
	c1 := cmd.New("abc", "does something", func(c *cmd.Command[c1Args]) error { fmt.Println(*c.Store().oye, *c.Store().someval); return nil })
	c1.Store().oye = c1.Bool("oye", "oyes around")
	c1.Store().someval = c1.Int("int", 1, "some int")

	c2 := cmd.New("oyebc", "does something else", func(c *cmd.Command[c1Args]) error { fmt.Println(*c.Store().oye, *c.Store().someval); return nil })
	c2.Store().oye = c2.Bool("oyeof2", "oyes around2")
	c2.Store().someval = c2.Int("intfor2", 123, "some int2")

	w2 := cmd.NewWrapper("sub", "subsub test")
	w2.AddCommand(c1)
	w2.AddCommand(c2)

	w3 := cmd.NewWrapper("subsub", "subsubsubsubsubs")
	w3.AddCommand(w2)

	w1 := cmd.NewWrapper("root", "is the root executor")
	w1.AddCommand(w3)
	w1.AddCommand(w2)

	err := w1.ParseOsArgsAndExec()
	fmt.Println(err)
}
