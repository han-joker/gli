package cmd

import (
	"flag"
	"fmt"
)

type GLI struct {
	commands map[string]CMD
}

func New() *GLI {
	return &GLI {
		commands: map[string]CMD{},
	}
}

func (g *GLI) Run() {
	g.initCmd()
	flag.Parse()
	g.Usage()
}

func (g *GLI) initCmd() {
	versionCMD := &Version{}
	versionCMD.Title = "version"
	versionCMD.FlagSet = flag.NewFlagSet("version", flag.ExitOnError)
	g.commands["version"] = versionCMD
}

func (g *GLI) Usage() {
	fmt.Println("Gli is a Scaffold based on gin.")
	fmt.Println("Version: 0.0.0")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("\tgli <command> [arguments]")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println()
	fmt.Printf("\t%s\t\t%s\n", "build", "build code.")
	fmt.Printf("\t%s\t\t%s\n", "init", "init a gli app based on go module.")
	fmt.Printf("\t%s\t\t%s\n", "version", "print Go version")
	fmt.Println()
	fmt.Println("Use `go help <command>` for more information about a command.")
}