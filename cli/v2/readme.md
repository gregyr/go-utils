# CLI Application Package

## Commands

Commands are the actual building blocks of the CLI Application, they are the ones that execute functions.

### Creating a Command
```go
cmd := cli.New("command", "description of this command", func(c *cli.Command[CommandArgs]) error {
    fmt.Printf("arg1: %v, arg2: %v\n", *c.Store().arg1, *c.Store().arg2)
    return nil
})
```
This creates a new command named `command` with a description and a function that is called when the command is used. It is of type CommandArgs, that is the type of its Store, where Flag values are stored in, that can be a simple struct:

```go
type CommandArgs struct {
	arg1 *bool
	arg2 *string
}
```
To actually add the flags, you have to register them through the command object and store them into the commands store:
```go
cmd.Store().arg1 = cmd.Bool("arg1", "some boolean argument")
cmd.Store().arg2 = cmd.String("arg2", "some string argument", "abc")

// this would be registered but not stored
cmd.Int("arg3", "some integer argument", 0)
```
To Run your newly created command Parse the OS's Args and Execute the command:
```go
cmd.ParseOsArgs()
cmd.Exec()
```
### Options
#### Explanations
```go
cmd.AddExplanation("This command expects a certain amount of arguments to work, otherwise it will throw an error. it also has an explicit usage string instead of the generic one, cool right?", 80)
```
You can add a longer explanation with a desired maximum width when displayed in the usage info. If the maximum width is set to 0, no line wrapping is applied

#### Usage String
```go
cmd.ExplicitUsageString("<name> [flags] [from] [to]")
```
This sets the usage string displayed in the usage info, replacing the generic version

## Wrappers

Wrappers can be used to manage subcommands. They can be seen like folders in a file system.

### Creating a Wrapper
```go
w := cli.NewWrapper("abc", "this is a command description")
```
This creates a Wrapper with the command name `abc`. It can hold and organize multiple subcommands
```go
w.AddCommand(cmd)
```
This adds our previously created command as a subcommand of this wrapper, you can add as many commands as you want.
```go
w.ParseOsArgsAndExec()
```
This runs the wrapper by parsing the os arguments and executing the wrapper command as the root command. From here it steps through its subcommands until it hits an actual command and executes it.

## Example Application

```go

// a struct to manage flags for a command
type CommandArgs struct {
	arg1 *bool
	arg2 *string
}

// create your root wrapper
w := cli.NewWrapper("abc", "this is a command description")

// add an optional explanation with a max width of 80 characters per line
w.AddExplanation("This is meant to be a longer section in the help of a command to explain in detail what the command is actually meant to do. This can contain a description of licenses, usage guides, the idea behind a command or anything else. \nThis section will be rendered seperately just underneath the short description of the command. having this is optional.", 80)

// create a command
cmd := cli.New("command", "description of this command", func(c *cli.Command[CommandArgs]) error {
    fmt.Printf("arg1: %v, arg2: %v\n", *c.Store().arg1, *c.Store().arg2)
    return nil
})
// register the flags for a command and their aliases
cmd.Store().arg1 = cmd.Bool("arg1", "some boolean argument")
cmd.Store().arg2 = cmd.String("arg2", "some string argument", "abc")
cmd.Alias("a", "arg1")
cmd.Alias("argument1", "arg1")
cmd.Alias("b", "arg2")
cmd.Alias("argument2", "arg2")
// these flags are not stored, so you wouldn't be able to access them
cmd.Bool("test", "a testing command")
cmd.Bool("somethingelse", "another flag")
cmd.Bool("more", "i dont know anymore")
// add the command to your wrapper
w.AddCommand(cmd)

// add another command
cmd2 := cli.New("anotherone", "description of this other command", func(c *cli.Command[any]) error { fmt.Printf("tail: %v\n", c.Tail()); return nil })
// this command expects exactly two arguments in its Tail
cmd2.SetTailMode(cli.TailModeExact, 2) 
// this is how it should be used
cmd2.ExplicitUsageString("<name> [flags] [from] [to]")
cmd2.Bool("flag", "this is a cool falg")
// this is a longer explanaition of what it does, the explanation has a max width of 80 characters
cmd2.AddExplanation("This command expects a certain amount of arguments to work, otherwise it will throw an error. it also has an explicit usage string instead of the generic one, cool right?", 80)

// add some more commands
w.AddCommand(cli.New("short", "shorter command with a short description", func(c *cli.Command[any]) error { return nil }))
w.AddCommand(cli.New("longer-command-123", "description of a really long command with a long description", func(c *cli.Command[any]) error { return nil }))

// run the commands
w.ParseOsArgsAndExec()
```

