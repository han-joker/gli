package cmd

import (
	"fmt"
	"github.com/han-joker/gli/gen"
	"os"
)

type Gen struct {
	baseCMD
}

func (c *Gen) Run() {
	//c.FlagSet.Usage = c.Usage
	c.FlagSet.Parse(os.Args[2:])

	if c.FlagSet.NArg() == 0 {
		c.Usage()
		return
	}

	switch c.FlagSet.Arg(0) {
	case "router":
		if c.FlagSet.NArg() == 1 {
			c.RouterUsage()
			return
		}
		for i:=1; i<c.FlagSet.NArg(); i++ {
			gen.GenRouterResource(c.FlagSet.Arg(i))
		}
	default:
		c.Usage()
		return
	}

}

func (c *Gen) Usage() {
	fmt.Println("gli gen provides generate code.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Printf("\t%s\n", "gli gen api resource")
	fmt.Printf("\t%s\n", "gli gen router resource")
	fmt.Printf("\t%s\n", "gli gen handler resource")
	//fmt.Printf("\t%s\n", "gli gen model resource")
	fmt.Printf("\t%s\n", "gli gen config key default")
	fmt.Println()
	fmt.Printf("\t%s\n", "api = router + handler")

}

func (c *Gen) RouterUsage() {
	fmt.Println("Usage:")
	fmt.Printf("\t%s\n", "gli gen router <resource>[ resource]")
	fmt.Printf("\n")
	fmt.Println("Examples:")
	fmt.Printf("\t%s\n", "gli gen router product")
}

