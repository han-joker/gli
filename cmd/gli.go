package cmd

import (
	"flag"
	"fmt"
	"os"
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

	cmdStr := ""
	if len(os.Args) > 1 {
		cmdStr = os.Args[1]
	}
	if cmd, ok := g.commands[cmdStr]; ok {
		cmd.Run()
	} else {
		g.Usage()
	}
}

func (g *GLI) initCmd() {
	//g.initCMDVersion()
	g.initCMDInit()
}

func (g *GLI) initCMDInit() {
	cmd := &Init{}
	cmd.Title = "init"
	cmd.FlagSet = flag.NewFlagSet("init", flag.ExitOnError)
	g.commands["init"] = cmd
}

//func (g *GLI) initCMDVersion() {
//	cmd := &Version{}
//	cmd.Title = "version"
//	cmd.FlagSet = flag.NewFlagSet("version", flag.ExitOnError)
//	g.commands["version"] = cmd
//}

func (g *GLI) Usage() {
	fmt.Println("Gli is a Scaffold based on gin.")
	fmt.Println("Version: 0.0.0")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("\tgli <command> [arguments]")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println()
	//fmt.Printf("\t%s\t\t%s\n", "build", "build router, handler and so on.")
	fmt.Printf("\t%s\t\t%s\n", "init", "init a gin application based on go module.")
	//fmt.Printf("\t%s\t\t%s\n", "version", "print Go version")
	fmt.Println()
	fmt.Println("Use `gli help <command>` for more information about a command.")
}