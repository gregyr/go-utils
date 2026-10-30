package cli

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gregyr/go-utils/slice"
)

var (
	fileNameIsDisplayName bool = false
)

type ArgParser interface {
	ParseAndExec([]string, bool) error
	Name() string
	Desc() string
}

type Wrapper struct {
	Formatting
	subCommands map[string]ArgParser
}

// creates a new command Wrapper which can hold sub commands
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

// parses the given arguments and executes the command, if isroot is true, the command is handled as the entrypoint of the applications, this may affect things like the usage info
func (w *Wrapper) ParseAndExec(args []string, isroot bool) error {
	var cmdname string
	if isroot && fileNameIsDisplayName {
		cmdname = filepath.Base(args[0])
	} else {
		cmdname = w.name
	}
	args = args[1:]

	if len(args) == 0 {
		w.usage(isroot)
		fmt.Println()
		return fmt.Errorf("%s - expected subcommand\n", cmdname)
	}

	if args[0] == "-h" || args[0] == "--help" || args[0] == "help" {
		w.usage(isroot)
		return nil
	}

	nextcmd, ok := w.subCommands[args[0]]
	if !ok {
		w.usage(isroot)
		fmt.Println()
		return fmt.Errorf("%s - unknown subcommand %s\n", cmdname, args[0])
	}
	return nextcmd.ParseAndExec(args, false)
}

// Runs the cli application by automatically parsing the arguments and executing the command. If errors occur it exits with the usage and error message
func (w *Wrapper) ParseOsArgsAndExec() {
	args := os.Args

	err := w.ParseAndExec(args, true)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func (w *Wrapper) usage(isroot bool) {
	var b strings.Builder

	w.writeHeader(&b, w.genericUsageString(isroot))

	fmt.Fprintf(&b, "Available Commands are:\n")
	cmdNames := slice.IterToSlice(maps.Keys(w.subCommands))

	slices.Sort(cmdNames)
	maxLen := 0
	for _, n := range cmdNames {
		maxLen = max(maxLen, len(n))
	}
	padding := max(maxLen, 12)
	for _, n := range cmdNames {
		cmd := w.subCommands[n]
		formatString := fmt.Sprintf("\t%%-%vs %%s\n", padding)
		fmt.Fprintf(&b, formatString, cmd.Name(), cmd.Desc())
	}

	fmt.Fprintf(&b, "\nUse \"%s [-h | --help | help]\" to show this menu\n", w.name)
	fmt.Fprintf(&b, "or use \"%s <command> [-h | --help]\" for more information about a command\n", w.name)

	fmt.Print(b.String())
}

func (w *Wrapper) genericUsageString(isroot bool) func() string {
	return func() string {
		var programName string = w.name
		if isroot && fileNameIsDisplayName {
			programName = filepath.Base(os.Args[0])
		}

		return fmt.Sprintf("%s <command> [arguments]\n\n", programName)
	}
}
