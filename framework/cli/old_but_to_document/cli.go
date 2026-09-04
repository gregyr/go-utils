package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/gregyr/go-utils/util/core/fun"
)

type Flag struct {
	short string
	long  string
	info  string
}

type Opt struct {
	short string
	long  string
	info  string
}

type CLI struct {
	name       string
	inputName  string
	input      string
	flags      []Flag
	opts       []Opt
	flagstates []bool
	optsVals   []string
}

// creates a new CLI object
func New(name, inputName string) CLI {
	return CLI{flags: make([]Flag, 0), name: name, inputName: inputName}
}

// Creates a new Flag
func NewFlag(short, long, info string) Flag {
	return Flag{short: short, long: long, info: info}
}

// Creates a new Opt
func NewOpt(short, long, info string) Opt {
	return Opt{short: short, long: long, info: info}
}

// adds a flag parameter, the parser checks if flags are provided or not but don't associate any value with these, for inputs use an opt
func (f *CLI) AddFlag(flag Flag) {
	f.flags = append(f.flags, flag)
}

// adds an option parameter, meaning the parser expects an input value after this
func (c *CLI) AddOpt(opt Opt) {
	c.opts = append(c.opts, opt)
}

// Prints the usage overview
func (f *CLI) Usage() {

	var builder strings.Builder

	builder.WriteString("Usage:\n")

	fmt.Fprintf(&builder, "%s %s %s\n\n", f.name, fun.If(len(f.flags) > 0, "[options/flags]", ""), fun.If(len(f.inputName) > 0, fmt.Sprintf("[%s]", f.inputName), ""))

	builder.WriteString("Options:\n")
	for _, opt := range f.opts {
		fmt.Fprintf(&builder, "  %-20s   %s\n", fmt.Sprintf("--%s, -%s", opt.long, opt.short), opt.info)
	}

	builder.WriteString("Flags:\n")
	for _, flag := range f.flags {
		fmt.Fprintf(&builder, "  %-20s   %s\n", fmt.Sprintf("--%s, -%s", flag.long, flag.short), flag.info)
	}

	fmt.Println(builder.String())
}

func (f *CLI) getFlagIdxShort(flagchar rune) int {
	for i, flag := range f.flags {
		if flag.short == string(flagchar) {
			return i
		}
	}
	return -1
}

func (f *CLI) getFlagIdxLong(flagstr string) int {
	for i, flag := range f.flags {
		if flag.long == flagstr {
			return i
		}
	}
	return -1
}

func (f *CLI) getOptIdxShort(optchar byte) int {
	for i, opt := range f.opts {
		if opt.short == string(optchar) {
			return i
		}
	}
	return -1
}

func (f *CLI) getOptIdxLong(optstr string) int {
	for i, opt := range f.opts {
		if opt.long == optstr {
			return i
		}
	}
	return -1
}

// returns immediately if -h or --help is provided, if an invalid argument is provided an error is returned
func (f *CLI) LoadArgs() error {

	args := os.Args

	return f.parseArgs(args)
}

func (f *CLI) parseArgs(args []string) error {
	optIdx := 0
	isOptValue := false

	f.flagstates = make([]bool, len(f.flags))
	f.optsVals = make([]string, len(f.opts))

	if len(args) == 0 {
		return nil
	}

	for i, arg := range args[1:] {

		if isOptValue {
			if len(arg) <= 0 {
				return fmt.Errorf("Missing Value for Option")
			}
			f.optsVals[optIdx] = arg
			isOptValue = false
			continue
		}

		if len(arg) <= 0 {
			continue
		}

		// handle both long options and flags, since it's max one at a time for either
		if len(arg) > 2 && arg[0:2] == "--" {

			// display help usage message
			if arg[2:] == "help" {
				f.Usage()
				os.Exit(0)
			}

			idx := f.getFlagIdxLong(arg[2:])
			optIdx := f.getOptIdxLong(arg[2:])

			if idx != -1 {
				f.flagstates[idx] = true
			} else if optIdx != -1 {
				isOptValue = true
				optIdx = idx
			} else {
				return fmt.Errorf("Unknown arg: %s", arg[2:])
			}

			continue
		}

		// handle option short variant, if not an option continues into next if to handle short versions of flags
		if len(arg) == 2 && arg[0] == '-' {

			// display help usage message
			if arg[1] == 'h' {
				f.Usage()
				os.Exit(0)
			}

			char := arg[1]
			idx := f.getOptIdxShort(char)

			if idx != -1 {
				isOptValue = true
				optIdx = idx
				continue
			}
		}

		// handle short version of flags
		if len(arg) >= 2 && arg[0] == '-' {

			// display help usage message
			if arg[1] == 'h' {
				f.Usage()
				os.Exit(0)
			}

			for _, char := range arg[1:] {
				idx := f.getFlagIdxShort(char)
				if idx != -1 {
					f.flagstates[idx] = true
				} else {
					return fmt.Errorf("Unknown arg: %s", string(char))
				}
			}

			continue
		}

		if i == len(args)-2 && len(arg) > 0 && arg[0] != '-' {
			f.input = arg
		}
	}
	return nil
}

// returns whether a flag has been provided, returns false when the flag is invalid aswell
func (c *CLI) Has(flag string) bool {
	for i, f := range c.flags {
		if (f.short == flag || f.long == flag) && c.flagstates[i] {
			return true
		}
	}
	return false
}

// gets the input to the application, this is the last argument in the call, that is not a flag, an option or an option Value
func (c *CLI) GetInput() string {
	return c.input
}

// gets the input value associated with an option key, if the option was not provided or the key is invalid, an empty string is returned
func (c *CLI) GetOption(key string) string {
	for i, opt := range c.opts {
		if opt.short == key || opt.long == key {
			return c.optsVals[i]
		}
	}
	return ""
}
