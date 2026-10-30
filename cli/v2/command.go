// package cli provides a basic framework for creating commands containing flags. Flags can be parsed via a CLI object. Supports both long and short flag identifiers aswell as aliases. Allows for chaining of short identifiers
package cli

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/gregyr/go-utils/slice"
)

// CLI IMPLEMENTATION
// what would be nice to have:
// input values seperated by space: --port 8080
// optionally parse flags after input: do this.txt --verbose

const (
	FlagBool string = "Bool"
)

const (
	TailModeMin TailMode = iota
	TailModeMax
	TailModeExact
)

type TailMode int

type TailSettings struct {
	mode   TailMode
	length int
}

type FlagValue interface {
	Set(string) error
	String() string
	Type() string
}

type Flag struct {
	Value FlagValue
	Desc  string
}

// the command type is used for executing a function when called. it is provided with an argstorage, this is usually a struct containing the values of flags if they are set.
type Command[T any] struct {
	Formatting
	flags        map[string]Flag
	aliases      map[string]string
	tail         []string
	tailSettings TailSettings
	fn           func(*Command[T]) error
	argstorage   *T
}

type stringValue string

func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}
func (s *stringValue) String() string { return string(*s) }
func (s *stringValue) Type() string   { return "String" }

type intValue int

func (i *intValue) Set(val string) error {
	v, err := strconv.Atoi(val)
	if err != nil {
		return fmt.Errorf("invalid integer value: %w", err)
	}
	*i = intValue(v)
	return nil
}
func (i *intValue) String() string { return strconv.Itoa(int(*i)) }
func (i *intValue) Type() string   { return "Int" }

type boolValue bool

func (b *boolValue) Set(val string) error {
	v, err := strconv.ParseBool(val)
	if err != nil {
		return fmt.Errorf("invalid boolean value: %w", err)
	}
	*b = boolValue(v)
	return nil
}
func (b *boolValue) String() string { return strconv.FormatBool(bool(*b)) }
func (b *boolValue) Type() string   { return FlagBool }

type float32Value float32

func (f *float32Value) Set(val string) error {
	v, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return fmt.Errorf("invalid float32 value: %w", err)
	}
	*f = float32Value(v)
	return nil
}
func (f *float32Value) String() string { return strconv.FormatFloat(float64(*f), 'f', -1, 32) }
func (f *float32Value) Type() string   { return "Float32" }

// Creates a new CLI object intended for single use, if it parsed args once, it cannot reset values to their defaults for subsequent calls
func New[T any](name, desc string, fn func(*Command[T]) error) *Command[T] {
	return &Command[T]{name: name, desc: desc, flags: make(map[string]Flag), aliases: make(map[string]string), fn: fn, argstorage: new(T), tailSettings: TailSettings{mode: TailModeMin, length: 0}}
}

// Creates a flag containing a boolean Value if named "help" or "h", the flag will not be found as those are reserved to the usage reference, the flag is false by default
func (c *Command[T]) Bool(identifier string, desc string) *bool {
	if err := errorOnInvalidIdentifier(identifier); err != nil {
		panic(err)
	}
	c.panicOnDoubleAssignment(identifier)

	p := new(bool)
	*p = false
	c.flags[identifier] = Flag{Value: (*boolValue)(p), Desc: desc}
	return p
}

// Creates a flag containing a int Value if named "help" or "h", the flag will not be found as those are reserved to the usage reference
func (c *Command[T]) Int(identifier string, def int, desc string) *int {
	if err := errorOnInvalidIdentifier(identifier); err != nil {
		panic(err)
	}
	c.panicOnDoubleAssignment(identifier)

	p := new(int)
	*p = def
	c.flags[identifier] = Flag{Value: (*intValue)(p), Desc: desc}
	return p
}

// Creates a flag containing a string Value if named "help" or "h", the flag will not be found as those are reserved to the usage reference
func (c *Command[T]) String(identifier string, def string, desc string) *string {
	if err := errorOnInvalidIdentifier(identifier); err != nil {
		panic(err)
	}
	c.panicOnDoubleAssignment(identifier)

	p := new(string)
	*p = def
	c.flags[identifier] = Flag{Value: (*stringValue)(p), Desc: desc}
	return p
}

// Creates a flag containing a float32 Value if named "help" or "h", the flag will not be found as those are reserved to the usage reference
func (c *Command[T]) Float32(identifier string, def float32, desc string) *float32 {
	if err := errorOnInvalidIdentifier(identifier); err != nil {
		panic(err)
	}
	c.panicOnDoubleAssignment(identifier)

	p := new(float32)
	*p = def
	c.flags[identifier] = Flag{Value: (*float32Value)(p), Desc: desc}
	return p
}

// Creates an alias to an existing flag. Panics if a flag does not exist or the alias has been assigned to a different flag or alias
func (c *Command[T]) Alias(alias, original string) {
	if err := errorOnInvalidIdentifier(alias); err != nil {
		panic(err)
	}
	c.panicOnDoubleAssignment(alias)

	if _, referencesAlias := c.aliases[original]; referencesAlias {
		panic(fmt.Sprintf("alias must reference canonical names %s: %s is an alias", alias, original))
	}

	if _, exists := c.flags[original]; !exists {
		panic(fmt.Sprintf("creating alias %s: %s has not been declared", alias, original))
	}
	c.aliases[alias] = original
}

// provides the remaining arguments after and including the first non flag argument or after and excluding "--"
func (c *Command[T]) Tail() []string {
	return slices.Clone(c.tail)
}

// returns the commands argument store, flags need to be written to this store for the command to be able to use them
func (c *Command[T]) Store() *T {
	return c.argstorage
}

// determines how the command handles tail arguments and how many it expects
func (c *Command[T]) SetTailMode(mode TailMode, length int) {
	c.tailSettings = TailSettings{mode: mode, length: length}
}

// Reads from os.Args and parses any provided arguments. Flags will be stored into previously created Flags, if a flag was not found, the program exits with a usage reference based on the created flags.
// parseing intercepts any flag provided as "-h" or "--help" by exiting the application with a usage reference
// This is single use, if any parsing is called a second time, it does not reset previous arguments
// This executes the command as the root command, meaning the command name will be overridden by the program name
func (c *Command[T]) ParseOrExit() {
	args := os.Args

	err := c.parse(args)
	if err != nil {
		fmt.Println(err)
		c.usage()
		os.Exit(1)
	}
}

// Reads from os.Args and parses any provided arguments. Flags will be stored into previously created Flags, if a flag was not found, the program exits with an error.
// Parse intercepts any flag provided as "-h" or "--help" by exiting the application with a usage reference
// This is single use, if any parsing is called a second time, it does not reset previous arguments
// This executes the command as the root command, meaning the command name will be overridden by the program name
func (c *Command[T]) ParseOsArgs() error {
	args := os.Args

	return c.parse(args)
}

// parses the given args into the flag values. may produce partially malformed state if parseing fails halfway through
// This is single use, if parse is called a second time, it does not reset previous arguments
func (c *Command[T]) parse(args []string) error {
	if len(args) < 1 {
		c.tail = []string{}
		return nil
	}

	token := 1

	for token < len(args) {
		arg := args[token]

		// empty strings can cause length of 0, is not a flag so terminate flags
		if len(arg) == 0 {
			break
		}
		isFlag := arg[0] == '-'
		if !isFlag {
			break
		}

		flagContent := arg[1:]
		if len(flagContent) == 0 {
			return fmt.Errorf("missing flag after '-'")
		}

		// second minus -> long flag
		if flagContent[0] == '-' {
			actualContent := flagContent[1:]

			if endOfFlags := len(actualContent) == 0; endOfFlags {
				token++ // advance token to remove "--" from the tail
				break
			}

			flagname, value, found := strings.Cut(actualContent, "=")
			if flagname == "help" {
				if found != false {
					return fmt.Errorf("invalid usage of 'help'")
				}
				c.usage()
				os.Exit(0)
			}

			if len(flagname) < 2 {
				return fmt.Errorf("unknown flag --%s, did you mean -%s\n", flagname, flagname)
			}

			flag, err := c.getFlag(flagname)
			if err != nil {
				return err
			}

			err = c.setOrFlip(flag, flagname, value, found)
			if err != nil {
				return err
			}

		} else { // no second '-'
			flagname, value, found := strings.Cut(flagContent, "=")
			if flagname == "h" {
				if found != false {
					return fmt.Errorf("invalid usage of 'h'")
				}
				c.usage()
				os.Exit(0)
			}
			if found {
				if len(flagname) != 1 {
					return fmt.Errorf("unknown flag -%s, did you mean --%s\n", flagname, flagname)
				} else {
					flag, err := c.getFlag(flagname)
					if err != nil {
						return err
					}
					err = c.setOrFlip(flag, flagname, value, found)
					if err != nil {
						return err
					}
				}
			} else {
				for _, char := range flagContent {
					if char == 'h' {
						return fmt.Errorf("invalid usage of 'h'")
					}
					flag, err := c.getFlag(string(char))
					if err != nil {
						return err
					}
					err = c.setOrFlip(flag, string(char), "", false)
					if err != nil {
						return err
					}
				}
			}
		}

		token++

	}
	c.tail = slices.Clone(args[token:])

	tailMode := c.tailSettings.mode
	wantedTailLength := c.tailSettings.length
	tailLength := len(c.tail)

	if tailMode == TailModeExact && tailLength != wantedTailLength {
		return fmt.Errorf("Expected %v arguments, got %v", wantedTailLength, tailLength)
	}
	if tailMode == TailModeMin && tailLength < wantedTailLength {
		return fmt.Errorf("Expected atleast %v arguments, got %v", wantedTailLength, tailLength)
	}
	if tailMode == TailModeMax && tailLength > wantedTailLength {
		return fmt.Errorf("Expected at most %v arguments, got %v", wantedTailLength, tailLength)
	}

	return nil
}

func (c *Command[T]) Exec() error {
	return c.fn(c)
}

func (c *Command[T]) ParseAndExec(args []string, _ bool) error {
	err := c.parse(args)
	if err != nil {
		return err
	}
	return c.Exec()
}

// displays the usage of flags with their types
func (c *Command[T]) usage() {
	var b strings.Builder

	c.writeHeader(&b, c.genericUsageString)

	const UNDERLINE string = "\033[1;4m"
	const RESET string = "\033[0m"

	flagnames := slice.IterToSlice(maps.Keys(c.flags))
	sortFlagNames(flagnames, c.flags)
	fmt.Fprintf(&b, "Available Flags are:\n")
	for _, name := range flagnames {
		oneChar, long := getSplitFlagNames(c.aliasesFor(name))
		fmt.Fprintf(&b, "\t%s", UNDERLINE)
		for i, flag := range oneChar {
			fmt.Fprintf(&b, "-%s", flag)
			if i != len(oneChar)-1 || len(long) > 0 {
				b.WriteString(", ")
			}
		}

		for i, flag := range long {
			fmt.Fprintf(&b, "--%s", flag)
			if i != len(long)-1 {
				b.WriteString(", ")
			}
		}
		fmt.Fprintf(&b, "%s", RESET)

		fflag := c.flags[name]
		if fflag.Value.Type() == FlagBool {
			fmt.Fprintf(&b, ":\n\t    %s\n", c.flags[name].Desc)
		} else {
			fmt.Fprintf(&b, " = %v (%v):\n\t    %s\n", fflag.Value.Type(), fflag.Value, fflag.Desc)
		}
	}
	fmt.Print(b.String())
}

func (c *Command[T]) panicOnDoubleAssignment(identifier string) {
	if _, ok := c.flags[identifier]; ok {
		panic("Double assigned Flag identifier")
	}
	if _, ok := c.aliases[identifier]; ok {
		panic("Double assigned Flag identifier")
	}
}

// based on a root identifier: gets the aliases for it
func (c *Command[T]) aliasesFor(identifier string) []string {
	aliases := make([]string, 1)
	aliases[0] = identifier
	for a, f := range c.aliases {
		if f == identifier {
			aliases = append(aliases, a)
		}
	}
	return aliases
}

// gets the flag corresponding to the identifier, this can be the canonical or an alias
func (c *Command[T]) getFlag(identifier string) (Flag, error) {
	flag, ok := c.flags[identifier]
	if !ok {
		alias, aliasexists := c.aliases[identifier]
		if !aliasexists {
			return Flag{}, fmt.Errorf("unknown flag: %s", identifier)
		}
		flagfromalias, flagfromaliasexists := c.flags[alias]
		if !flagfromaliasexists {
			return Flag{}, fmt.Errorf("unknown flag (from alias): %s", alias)
		}
		flag = flagfromalias
	}

	return flag, nil
}

// sets any value based on the provided value string, if an error during setting occurs, it is bubbled up.
// in case of a boolean, if no value is present it is simply turned to true
func (c *Command[T]) setOrFlip(flag Flag, flagname, value string, found bool) error {
	// if len(value) == 0 && found == true {
	// 	return fmt.Errorf("missing value for flag %s", flagname)
	// }
	if flag.Value.Type() == FlagBool {
		if !found {
			flag.Value.Set("true")
		} else {
			err := flag.Value.Set(value)
			if err != nil {
				return fmt.Errorf("Invalid input for flag %s: %s", flagname, value)

			}
		}
	} else {
		if !found {
			return fmt.Errorf("flag %s requires an input like %s=%s", flagname, flagname, flag.Value)
		}
		err := flag.Value.Set(value)
		if err != nil {
			return fmt.Errorf("Invalid input for flag %s: %s", flagname, value)

		}
	}
	return nil
}

func (c *Command[T]) genericUsageString() string {
	programName := c.name

	flagStr := "[flags]"
	argStr := "[arguments]"
	if len(c.flags) == 0 {
		flagStr = ""
	}
	if (c.tailSettings.mode == TailModeMax || c.tailSettings.mode == TailModeExact) && c.tailSettings.length == 0 {
		argStr = ""
	}
	return fmt.Sprintf("%s %s %s \n", programName, flagStr, argStr)
}

// splits a list of flagnames into the ones which are one character long and those longer.
// returns the individual flags in lexicographical order
func getSplitFlagNames(allFlagNames []string) (oneChar []string, long []string) {
	slices.SortFunc(allFlagNames, func(a, b string) int { return len(a) - len(b) })

	for _, flag := range allFlagNames {
		if len(flag) == 1 {
			oneChar = append(oneChar, flag)
		} else {
			long = append(long, flag)
		}
	}

	slices.Sort(oneChar)
	slices.Sort(long)

	return oneChar, long
}

// validates an identifier string, return a proper output error message incase of an invalid identifier
func errorOnInvalidIdentifier(identifier string) error {
	if len(identifier) == 0 {
		return fmt.Errorf("flag identifier must be atleast 1 character long: %s", identifier)
	}
	if identifier[0] == '-' {
		return fmt.Errorf("flag identifier can not start with '-': %s", identifier)
	}
	if identifier == "help" || identifier == "h" {
		return fmt.Errorf("Cannot register 'help'/'h', these are reserverd for the usage reference")
	}
	if strings.Contains(identifier, "=") {
		return fmt.Errorf("flag identifier cannot contain '=': %s", identifier)
	}
	if strings.ContainsAny(identifier, " \t\r\n._\\/<>'#+*~,;:!\"§$%&()?`[{}]") {
		return fmt.Errorf("flag identifier cannot contain whitespaces: %s", identifier)
	}
	return nil
}

// sort the slice in place such that first all the boolean flags are in lexicographical order, then the rest in lexicographicalorder
func sortFlagNames(flagnames []string, flags map[string]Flag) {
	slices.SortFunc(flagnames, func(a, b string) int {
		if (flags[a].Value.Type() == flags[b].Value.Type() && flags[a].Value.Type() == FlagBool) ||
			(flags[a].Value.Type() == flags[b].Value.Type() && flags[a].Value.Type() != FlagBool) {
			if a < b {
				return -1
			}
			if a > b {
				return 1
			}
			return 0
		}
		if flags[a].Value.Type() == FlagBool {
			return -1
		}
		return 1
	})
}
