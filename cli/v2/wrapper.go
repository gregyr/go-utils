package cmd

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gregyr/go-utils/fn"
	"github.com/gregyr/go-utils/slice"
)

// POC, final product needs to be more resilient
// make usage use actual args and not os.Args
// improve handling errors and subsequent usage output
// add -h, --help and help options to wrapper

type ArgParser interface {
	ParseAndExec(args []string) error
	Name() string
	Desc() string
}

type Wrapper struct {
	name        string
	desc        string
	subCommands map[string]ArgParser
}

func NewWrapper(name, desc string) *Wrapper {
	return &Wrapper{name: name, desc: desc, subCommands: make(map[string]ArgParser, 0)}
}

// panics if an already existing command for this wrapper is assigned
func (w *Wrapper) AddCommand(command ArgParser) {
	if _, exists := w.subCommands[command.Name()]; exists {
		panic(fmt.Sprintf("The command %s already exists", command.Name()))
	}
	w.subCommands[command.Name()] = command
}

func (c *Wrapper) Name() string {
	return c.name
}

func (c *Wrapper) Desc() string {
	return c.desc
}

func (w *Wrapper) ParseAndExec(args []string) error {

	cmdname := filepath.Base(args[0])
	args = args[1:]

	if len(args) == 0 {
		w.Usage(false)
		fmt.Println()
		return fmt.Errorf("%s - expected subcommand\n", cmdname)
	}

	if args[0] == "-h" || args[0] == "--help" || args[0] == "help" {
		w.Usage(false)
		return nil
	}

	nextcmd, ok := w.subCommands[args[0]]
	if !ok {
		w.Usage(false)
		fmt.Println()
		return fmt.Errorf("%s - unknown subcommand %s\n", cmdname, args[0])
	}
	return nextcmd.ParseAndExec(args)
}

func (w *Wrapper) ParseOsArgsAndExec() {
	args := os.Args

	err := w.ParseAndExec(args)
	if err != nil {
		os.Exit(1)
	}
}

func (w *Wrapper) Usage(isroot bool) {
	var b strings.Builder

	programName := fn.If(isroot, filepath.Base(os.Args[0]), w.name)

	fmt.Fprintf(&b, "Usage of %s\n", programName)
	cmdNames := slice.IterToSlice(maps.Keys(w.subCommands))

	slices.Sort(cmdNames)

	for _, n := range cmdNames {
		cmd := w.subCommands[n]
		fmt.Fprintf(&b, "%s - %s\n", cmd.Name(), cmd.Desc())
	}

	fmt.Fprintln(&b, "\n-h, --help, help - show this menu")

	fmt.Print(b.String())
}
