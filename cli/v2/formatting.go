package cli

import (
	"fmt"
	"io"
	"strings"
)

// common data between wrapper and command
type Formatting struct {
	name        string
	desc        string
	usageString string
	explanation Explanation
}

// sets an explicit usage string to be displayed in the help menu, any occurences of <name> will be replace by the command name
// example:
//
//	w.ExplicitUsageString("<name> [flags] [from] [to]")
func (f *Formatting) ExplicitUsageString(usageString string) {
	f.usageString = strings.Trim(usageString, "\n\r ")
}

func (f *Formatting) Name() string {
	return f.name
}

func (f *Formatting) Desc() string {
	return f.desc
}

// adds a more extensive explanation to the command, maxWidth controlls after how many characters the lines wrap, set it to 0, to disable linewrap
func (f *Formatting) AddExplanation(expl string, maxWidth int) {
	f.explanation.explanation = expl
	f.explanation.maxWidth = maxWidth
}

func (f *Formatting) writeHeader(w io.Writer, genericUsageString func() string) {

	fmt.Fprintf(w, "%s\n\n", f.desc)

	f.explanation.writeTo(w)
	fmt.Fprint(w, "Usage:\n\t")
	if f.usageString == "" {
		fmt.Fprintln(w, genericUsageString())
	} else {
		fmt.Fprintf(w, "%s\n\n", strings.ReplaceAll(f.usageString, "<name>", f.name))
	}
}

type Explanation struct {
	explanation string
	maxWidth    int
}

func (e Explanation) writeTo(w io.Writer) {
	if e.explanation == "" {
		return
	}

	if e.maxWidth <= 0 {
		fmt.Fprintf(w, "%s\n", e.explanation)
		return
	}

	for _, line := range strings.Split(e.explanation, "\n") {
		for len(line) > 0 {
			if len(line) <= e.maxWidth {
				fmt.Fprintln(w, line)
				break
			}

			cut := e.maxWidth
			if idx := strings.LastIndexAny(line[:e.maxWidth], " \t"); idx > 0 {
				cut = idx
			}

			wrapped := strings.TrimRight(line[:cut], " \t")
			if wrapped == "" {
				wrapped = line[:e.maxWidth]
			}

			fmt.Fprintln(w, wrapped)
			line = strings.TrimLeft(line[cut:], " \t")
		}
	}

	fmt.Fprintln(w)
}
