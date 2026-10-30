package main

import (
	"fmt"

	"github.com/gregyr/go-utils/cli/v2"
)

type CommandArgs struct {
	arg1 *bool
	arg2 *string
}

func main() {
	w := cli.NewWrapper("abc", "this is a command description")

	w.ExplicitUsageString("<name> [command]")

	w.AddExplanation("This is meant to be a longer section in the help of a command to explain in detail what the command is actually meant to do. This can contain a description of licenses, usage guides, the idea behind a command or anything else. \nThis section will be rendered seperately just underneath the short description of the command. having this is optional.", 80)
	cmd := cli.New("command", "description of this command", func(c *cli.Command[CommandArgs]) error {
		fmt.Printf("arg1: %v, arg2: %v, tail: %v\n", *c.Store().arg1, *c.Store().arg2, c.Tail())
		return nil
	})
	cmd.Store().arg1 = cmd.Bool("arg1", "some boolean argument")
	cmd.Alias("a", "arg1")
	cmd.Alias("argument1", "arg1")
	cmd.Store().arg2 = cmd.String("arg2", "abc", "some string argument")
	cmd.Alias("b", "arg2")
	cmd.Alias("argument2", "arg2")
	cmd.Bool("oye", "oya")
	cmd.Bool("qwrfaf", "oya")
	cmd.Bool("fatte", "oya")
	w.AddCommand(cmd)
	cmd2 := cli.New("anotherone", "description of this other command", func(c *cli.Command[any]) error { fmt.Printf("tail: %v\n", c.Tail()); return nil })
	cmd2.SetTailMode(cli.TailModeExact, 2)
	cmd2.ExplicitUsageString("<name> [flags] [from] [to]")
	cmd2.Bool("flag", "this is a cool falg")
	cmd2.AddExplanation("This command expects a certain amount of arguments to work, otherwise it will throw an error. it also has an explicit usage string instead of the generic one, cool right?", 80)
	w.AddCommand(cmd2)
	w.AddCommand(cli.New("short", "shorter command with a short description", func(c *cli.Command[any]) error { return nil }))
	w.AddCommand(cli.New("longer-command-123", "description of a really long command with a long description", func(c *cli.Command[any]) error { return nil }))

	w.ParseOsArgsAndExec()
}
